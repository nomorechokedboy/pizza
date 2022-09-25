"use strict";
Object.defineProperties(exports, { __esModule: { value: true }, [Symbol.toStringTag]: { value: "Module" } });
var cookieParser = require("cookie-parser");
var cors = require("cors");
var express = require("express");
var http = require("http");
var morgan = require("morgan");
var mongoose = require("mongoose");
var slug = require("mongoose-slug-generator");
var bcrypt = require("bcrypt");
var jwt = require("jsonwebtoken");
var Validator = require("validatorjs");
var redis = require("redis");
function _interopDefaultLegacy(e) {
  return e && typeof e === "object" && "default" in e ? e : { "default": e };
}
var cookieParser__default = /* @__PURE__ */ _interopDefaultLegacy(cookieParser);
var cors__default = /* @__PURE__ */ _interopDefaultLegacy(cors);
var express__default = /* @__PURE__ */ _interopDefaultLegacy(express);
var morgan__default = /* @__PURE__ */ _interopDefaultLegacy(morgan);
var mongoose__default = /* @__PURE__ */ _interopDefaultLegacy(mongoose);
var slug__default = /* @__PURE__ */ _interopDefaultLegacy(slug);
var bcrypt__default = /* @__PURE__ */ _interopDefaultLegacy(bcrypt);
var jwt__default = /* @__PURE__ */ _interopDefaultLegacy(jwt);
var Validator__default = /* @__PURE__ */ _interopDefaultLegacy(Validator);
const serverError = "Server Error #";
const wrongUser = "Wrong email or password!";
const authFailed = "Authorized failed!";
mongoose__default["default"].plugin(slug__default["default"]);
const schema$1 = new mongoose__default["default"].Schema(
  {
    name: {
      type: String,
      require: true,
      unique: true
    },
    description: {
      type: String,
      required: true
    },
    img: {
      type: String,
      required: true,
      unique: true
    },
    price: {
      type: Number,
      required: true
    },
    type: {
      type: String,
      required: true,
      unique: true
    },
    slug: {
      type: String,
      slug: "name"
    }
  },
  { timestamps: true }
);
schema$1.index({ name: "text", description: "text" });
const productModel = mongoose__default["default"].model("product", schema$1);
const paging = async (req, res) => {
  const page = parseInt(req.query.page) || 0;
  const pageSize = parseInt(req.query.pageSize) || 20;
  const { search } = req.query;
  const cacheKey = `products:paging:page=${page}:pageSize=${pageSize}:search=${search}`;
  const client2 = req.app.get("redisClient");
  try {
    const hit = await client2.get(cacheKey);
    if (hit)
      return res.json({ data: JSON.parse(hit) });
    const products = await productModel.find(search ? { $text: { $search: search } } : {}).skip(page * pageSize).limit(pageSize).lean().exec();
    await client2.setEx(cacheKey, 3600, JSON.stringify(products));
    res.json({ data: products });
  } catch (e) {
    console.error(e);
    res.status(500).json({ error: serverError + new Date().toString() });
  }
};
const getDetail = async (req, res) => {
  const { slug: slug2 } = req.params;
  try {
    const hitRepo = await productModel.findOne({ slug: slug2 }).lean().exec();
    if (!hitRepo)
      return res.status(404).json({ error: `Cannot found product with slug ${slug2}` });
    res.json({ data: hitRepo });
  } catch (e) {
    console.error(e);
    res.status(500).json({ error: serverError + new Date().toString() });
  }
};
const getAll = async (req, res) => {
  const cacheKey = "products:getAll";
  const client2 = req.app.get("redisClient");
  try {
    const hit = await client2.get(cacheKey);
    res.cookie("refreshToken", "chanasdasdasdged", { maxAge: 60 * 1e3 * 5 });
    if (hit)
      return res.json({ data: JSON.parse(hit) });
    const products = await productModel.find({}).lean().exec();
    await client2.setEx(cacheKey, 60, JSON.stringify(products));
    res.json({ data: products });
  } catch (e) {
    console.log(e);
    res.status(500).json({ error: "Server Internal Error" });
  }
};
const postTest = (req, res, _) => {
  const { cookies } = req;
  console.log({ cookies });
  return res.json({ cookies });
};
const productRouter = express__default["default"].Router();
productRouter.post("/cookie", postTest);
productRouter.get("/paging", paging);
productRouter.get("/:slug/detail", getDetail);
productRouter.get("/", getAll);
const loginRules = {
  email: "required|email|min:8|max:30",
  password: ["required", "regex:/^(?=.*[a-z])(?=.*\\d)([^\\s]){8,}$/i"]
};
const {
  PIZZA_SERVER_PORT: PORT,
  MORGAN,
  MONGODB,
  SECRET_KEY,
  refreshTokenKey,
  REDIS_HOST,
  REDIS_PORT
} = process.env;
function genToken(payload) {
  if (!SECRET_KEY)
    throw Error(`Secret key is ${SECRET_KEY}`);
  return jwt__default["default"].sign(payload, SECRET_KEY, { expiresIn: "1m" });
}
const genRefreshToken = (payload) => {
  if (!refreshTokenKey)
    throw Error(`Refresh token key is ${refreshTokenKey}`);
  return jwt__default["default"].sign(payload, refreshTokenKey, { expiresIn: "2h" });
};
const schema = new mongoose__default["default"].Schema(
  {
    fullName: {
      type: String,
      required: true
    },
    email: {
      type: String,
      required: true,
      unique: true
    },
    phoneNumber: {
      type: String,
      required: true
    },
    password: {
      type: String,
      required: true
    },
    role: {
      type: String,
      required: true
    }
  },
  { timestamps: true }
);
schema.pre("save", function(next) {
  if (this.isNew) {
    const salt = bcrypt__default["default"].genSaltSync(10);
    this.password = bcrypt__default["default"].hashSync(this.password, salt);
  }
  next();
});
var User = mongoose__default["default"].model("user", schema);
const register$1 = async (req, res, next) => {
  const { email, password, phoneNumber, fullName } = req.body;
  const data = { email, password, fullName, phoneNumber };
  try {
    const find = await User.findOne({ email });
    if (find)
      return res.status(401).json({
        errors: [
          {
            email: [
              `The email ${email} has been used, please used another email!`
            ]
          }
        ]
      });
    const user = new User({
      ...data,
      role: "user"
    });
    const saveUser = await user.save();
    return res.status(201).json({ message: `Register success for email: ${saveUser.email}` });
  } catch (e) {
    next(e);
    res.status(500).json({ error: serverError + new Date().toString() });
  }
};
const login$1 = async (req, res, next) => {
  const { email, password } = req.body;
  try {
    const user = await User.findOne({ email });
    if (!user)
      return res.status(404).json({ error: wrongUser });
    const match = await bcrypt__default["default"].compare(password, user.password);
    if (!match)
      return res.status(404).json({ error: wrongUser });
    const { _id: id, role } = user;
    const payload = { id, role };
    return res.json({
      token: genToken(payload),
      refreshToken: genRefreshToken(payload)
    });
  } catch (e) {
    next(e);
    res.status(500).json({ error: serverError + new Date().toString() });
  }
};
const getUser = async (req, res, next) => {
  var _a;
  const { id } = req.params;
  const confirmId = (_a = req.decoded) == null ? void 0 : _a.id;
  try {
    if (id !== confirmId)
      return res.status(404).json({ error: "You have no authorized!" });
    const userInfo = await User.findById(id).select("fullName");
    if (!userInfo)
      return res.status(404).json({ error: "Invalid user!" });
    return res.json({ info: userInfo });
  } catch (e) {
    next(e);
    return res.status(500).json({ error: serverError + new Date().toString() });
  }
};
const refreshToken = (req, res, next) => {
  try {
    if (req.decoded) {
      const { id, role } = req.decoded;
      const payload = { id, role };
      res.json({
        token: genToken(payload),
        refreshToken: genRefreshToken(payload)
      });
    }
  } catch (e) {
    next(e);
    return res.status(500).json({ error: serverError + new Date().toString() });
  }
};
const register = (req, res, next) => {
  const registerRules = {
    ...loginRules,
    fullName: "required|min:8|max:30",
    phoneNumber: "required|min:10|max:12",
    cPassword: "same:password"
  };
  const validation = new Validator__default["default"](req.body, registerRules);
  if (validation.fails())
    return res.status(400).json(validation.errors.all());
  next();
};
const login = (req, res, next) => {
  const validation = new Validator__default["default"](req.body, loginRules);
  if (validation.fails())
    return res.status(400).json(validation.errors.all());
  next();
};
function isUser(req, res, next) {
  var _a;
  const token = (_a = req.headers.authorization) == null ? void 0 : _a.split(" ")[1];
  if (!token)
    return res.status(403).json({ error: "Invalid request (without token)" });
  if (!SECRET_KEY)
    throw Error(`Secret key is ${SECRET_KEY}`);
  try {
    const decoded = jwt__default["default"].verify(token, SECRET_KEY);
    req.decoded = decoded;
    if (decoded.role === "user")
      return next();
    throw Error(authFailed);
  } catch (e) {
    next(e);
    return res.status(401).json({ error: authFailed });
  }
}
function verifyToken(req, res, next) {
  var _a;
  const refreshToken2 = (_a = req.headers.authorization) == null ? void 0 : _a.split(" ")[1];
  if (!refreshToken2)
    throw Error("Refresh token is undefined!");
  if (!refreshTokenKey)
    throw Error("Refresh token key is undefined!");
  try {
    const decoded = jwt__default["default"].verify(refreshToken2, refreshTokenKey);
    if (decoded) {
      req.decoded = decoded;
      return next();
    }
    throw Error("Invalid token!");
  } catch (e) {
    next(e);
    return res.status(401).json({ error: "Invalid token!" });
  }
}
const userRouter = express__default["default"].Router();
userRouter.post("/register", register, register$1);
userRouter.post("/login", login, login$1);
userRouter.get("/refreshToken", verifyToken, refreshToken);
userRouter.get("/:id/detail", isUser, getUser);
const routerV1 = express.Router();
routerV1.use("/product", productRouter);
routerV1.use("/user", userRouter);
function connectDb() {
  if (!MONGODB)
    throw Error(`MONGODB is ${MONGODB}`);
  mongoose__default["default"].connect(MONGODB, {
    useNewUrlParser: true,
    useUnifiedTopology: true
  }).then(async () => {
    console.log("Connect Mongodb Success!");
  }).catch((e) => console.error(e));
}
const client = redis.createClient({
  socket: {
    host: REDIS_HOST || "localhost",
    port: REDIS_PORT || 6379
  }
});
client.on("error", (e) => console.log(`Redis client error: ${e}`));
const CORS_WHITELIST = [
  "http://localhost:5001/",
  "http://localhost:5000/",
  "https://pizza-api-nomorechokedboy.cloud.okteto.net"
];
const setup = async (app2) => {
  await client.connect();
  app2.set("redisClient", client);
  app2.use(cors__default["default"](CORS_WHITELIST));
  app2.use(express__default["default"].json({}));
  app2.use(cookieParser__default["default"]());
  app2.use(
    express__default["default"].urlencoded({
      extended: true
    })
  );
  morgan__default["default"].format(
    "myformat",
    '[:date[clf]] ":method :url" :status :res[content-length] - :response-time ms'
  );
  if (MORGAN === "1") {
    app2.use("/api/*", morgan__default["default"]("myformat"));
  }
  app2.use("/api/v1", routerV1);
  app2.use("/healthcheck", (_, res) => {
    const healthcheck = {
      uptime: process.uptime(),
      message: "I am fine!!!",
      timestamp: Date.now()
    };
    try {
      res.json(healthcheck);
    } catch (e) {
      console.log(e);
    }
  });
};
const app = express__default["default"]();
setup(app);
connectDb();
{
  const server = http.createServer(app);
  server.on("error", (e) => {
    if (e)
      throw e;
  });
  server.listen(PORT || 5e3, () => {
    console.log(`Pizza api on http://localhost:${PORT}`);
  });
}
const pizzaApi = app;
exports.pizzaApi = pizzaApi;
