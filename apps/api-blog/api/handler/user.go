package handler

import (
	"api-blog/api/presenter"
	"api-blog/api/util"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	usecase      usecase.UserUsecase
	authencation struct {
		JwtSecret       string
		JWTRefreshToken string
	}
}

func NewUserHandler(usecase usecase.UserUsecase, jwtSecret string, jwtExpiration string) *UserHandler {
	jwt := new(UserHandler)
	jwt.authencation.JwtSecret = jwtSecret
	jwt.authencation.JWTRefreshToken = jwtExpiration
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

	accessToken, refreshToken := util.GenerateToken(user.Identifier, []byte(handler.authencation.JwtSecret), []byte(handler.authencation.JWTRefreshToken))
	accessCookie, refreshCookie := util.GetAuthCookies(accessToken, refreshToken)
	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(presenter.Response{
		Status:  http.StatusOK,
		Message: "Login Success",
		Data: &fiber.Map{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"user":          user,
		},
	})
}

func (handler *UserHandler) GetUserById(c *fiber.Ctx) error {
	user := c.Locals("identifier")
	return c.JSON(presenter.Response{
		Status:  200,
		Message: "Nguoi day",
		Data: &fiber.Map{
			"user": user,
		},
	})
}
