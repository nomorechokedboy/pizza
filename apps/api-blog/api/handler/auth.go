package handler

import (
	"time"

	"api-blog/api/util"
	"api-blog/pkg/entities"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	JwtSecret       string
	JwtRefreshToken string
}

func NewAuthHanlder(jwtSecret string, jwtRefreshSecret string) *AuthHandler {
	return &AuthHandler{
		JwtSecret:       jwtSecret,
		JwtRefreshToken: jwtRefreshSecret,
	}
}

// GetNewAccessToken method for create a new access token.
// @Description Create a new access token.
// @Summary create a new access token
// @Tags Auth
// @Produce json
// @Param request body handler.RefreshToken.request true "refresh_token"
// @Success 200 {object} entities.Auth
// @Router /auth/refresh-token [post]
func (handler *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	type request struct {
		Refresh_token string `json:"refresh_token"`
	}
	tokenString := new(request)
	if err := c.BodyParser(tokenString); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	uId, err := util.ParseToken(tokenString.Refresh_token, []byte(handler.JwtRefreshToken))
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, err.Error())
	}
	accessToken := util.GenerateAccessClaims(uId, []byte(handler.JwtSecret), 15*time.Minute)
	newToken := new(entities.Auth)
	newToken.Token = accessToken
	newToken.RefreshToken = tokenString.Refresh_token
	return c.JSON(newToken)
}
