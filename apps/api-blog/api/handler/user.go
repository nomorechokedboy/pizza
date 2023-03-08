package handler

import (
	"api-blog/api/presenter"
	"api-blog/api/util"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	usecase      usecase.UserUsecase
	authencation struct {
		JwtSecret       string
		JWTRefreshToken string
	}
}

func NewUserHandler(usecase usecase.UserUsecase, jwtSecret string, jwtRefreshToken string) *UserHandler {
	jwt := new(UserHandler)
	jwt.authencation.JwtSecret = jwtSecret
	jwt.authencation.JWTRefreshToken = jwtRefreshToken
	jwt.usecase = usecase
	return jwt
}

// @CreateUser godoc
// @Summary Create User
// @Description Create New UserUsecase
// @Tags Users
// @Param todo body entities.UserReq true "New User"
// @Accept json
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 409
// @Router /api/v1/users/register [post]
func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	req := new(entities.UserReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	if _, err := handler.usecase.GetUserByIdentifier(req.Identifier); err == nil {
		return fiber.NewError(fiber.StatusConflict, "indentifier already existed")
	}
	if _, err := handler.usecase.GetUserByUsername(req.Username); err == nil {
		return fiber.NewError(fiber.StatusConflict, "username already existed")
	}
	err := handler.usecase.CreateUser(req.Password, req.Username, req.Fullname, req.PhoneNumber, req.Email, req.Avatar, req.Identifier)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create New user")
	}
	return c.JSON(presenter.Response{
		Status:  http.StatusCreated,
		Message: "Create success",
		Data:    nil,
	})
}

// Login
// @Login godoc
// @Summary User Login
// @Description Use for login response the refresh_token and accessToken
// @Tags Users
// @Accept json
// @Param todo body entities.UserLogin true "Login"
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 401
// @Failure 403
// @Router /api/v1/users/login [post]
func (handler *UserHandler) Login(c *fiber.Ctx) error {

	req := new(entities.UserLogin)
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

	accessToken, refreshToken := util.GenerateToken(user.Id, []byte(handler.authencation.JwtSecret), []byte(handler.authencation.JWTRefreshToken))

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

// GetAuthUserByToken
// @Login godoc
// @Summary Get user infor by token
// @Description Get UserInfo by accessToken
// @Tags Users
// @Accept json
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/users/ [get]
func (handler *UserHandler) GetAuthUserById(c *fiber.Ctx) error {
	authId, ok := c.Locals("uId").(uint)
	if !ok {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "can not parse Id from token")
	}
	user, err := handler.usecase.GetUserById(authId)
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(presenter.Response{
		Status:  200,
		Message: "Success",
		Data: &fiber.Map{
			"user": user,
		},
	})
}

// UpdateUserByToken
// @Update godoc
// @Summary Update user infor by token
// @Description Update UserInfo by Id from accessToken
// @Tags Users
// @Accept json
// @Param todo body entities.UserReq true "Updated User"
// @Success 200 {object} presenter.Response
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /api/v1/users/update [post]
func (handler *UserHandler) UpdateUserById(c *fiber.Ctx) error {
	authId, ok := c.Locals("uId").(uint)
	if !ok {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "can not parse Id from token")
	}
	user, err := handler.usecase.GetUserById(authId)
	if err != nil {
		return fiber.ErrNotFound
	}
	req := new(entities.UserReq)
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}
	if err := handler.usecase.UpdateUserInfo(req.Password, req.Fullname, req.PhoneNumber, req.Email, req.Avatar, user.Id); err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "can not update")
	}
	newUser, err := handler.usecase.GetUserById(user.Id)

	return c.JSON(presenter.Response{
		Status:  http.StatusOK,
		Message: "update success",
		Data: &fiber.Map{
			"user": newUser,
		},
	})
}

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Produce json
// @Success 200 {object} string
// @Security ApiKeyAuth
// @Router /api/v1/users/refresh_access_token [get]
func (handler *UserHandler) RefreshAccessToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(handler.authencation.JWTRefreshToken), nil
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	expirationtime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationtime) {
		return fiber.NewError(fiber.StatusUnauthorized, "token is out of date")
	}
	uId := uint(claims["sub"].(float64))
	accessToken := util.GenerateAccessClaims(uId, []byte(handler.authencation.JwtSecret), 15*time.Minute)
	return c.JSON(accessToken)
}
