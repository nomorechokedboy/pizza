"use strict";
Object.defineProperties(exports, { __esModule: { value: true }, [Symbol.toStringTag]: { value: "Module" } });
var express = require("express");
function _interopDefaultLegacy(e) {
  return e && typeof e === "object" && "default" in e ? e : { "default": e };
}
var express__default = /* @__PURE__ */ _interopDefaultLegacy(express);
const app = express__default["default"]();
app.get("/", (req, res) => {
  res.send("Updates, express~");
});
app.get("/ip", async (req, res) => {
  const resp = await fetch("https://api.ipify.org?format=json");
  const json = await resp.json();
  res.json(json);
});
{
  app.listen(3001);
  console.log("listening on http://localhost:3000/");
}
const viteNodeApp = app;
exports.viteNodeApp = viteNodeApp;
