package handler

import (
	"api-blog/api/presenter"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	usecase      usecase.UserUsecase
	authencation struct {
		JwtSecret     string
		JwtExpiration int64
	}
}

func NewUserHandler(usecase usecase.UserUsecase, jwtSecret string, jwtExpiration int64) *UserHandler {
	jwt := new(UserHandler)
	jwt.authencation.JwtSecret = jwtSecret
	jwt.authencation.JwtExpiration = jwtExpiration
	jwt.usecase = usecase
	return jwt
}

// @CreateUser godoc
// @Summary Create User
// @Description Create New UserUsecase
// @Tags Users
// @Accept json
// @Success 200
// @Failure 400
// @Router /register [post]
func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {

	req := new(entities.UserReq)

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	if _, err := handler.usecase.GetUserByIdentifier(req.Identifier); err == nil {
		return fiber.NewError(fiber.StatusConflict, "indentifier already existed")
	}
	if _, err := handler.usecase.GetUserByIdentifier(req.Username); err == nil {
		return fiber.NewError(fiber.StatusConflict, "username already existed")
	}
	err := handler.usecase.CreateUser(req.Password, req.Username, req.Fullname, req.PhoneNumber, req.Email, req.Avatar, req.Identifier)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create New user")
	}
	return nil
}

func (handler *UserHandler) Login(c *fiber.Ctx) error {
	type userLogin struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}
	req := new(userLogin)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	user, err := handler.usecase.GetUserByIdentifier(req.Identifier)
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "Email does not exist")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return fiber.NewError(fiber.StatusForbidden, "incorrect password")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Identifier
	claims["exp"] = handler.authencation.JwtExpiration
	tokenString, err := token.SignedString([]byte(handler.authencation.JwtSecret))

	// 	Value:    tokenString,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	HTTPOnly: true,
	// }
	// c.Cookie(&cookie)

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(presenter.Response{
		Status:  http.StatusOK,
		Message: "Login Success",
		Data: &fiber.Map{
			"token": tokenString,
			"user":  user,
		},
	})

}
