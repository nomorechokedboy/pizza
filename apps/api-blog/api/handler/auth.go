package handler

import (
	"api-blog/api/config"
	"api-blog/api/util"
	"api-blog/pkg/entities"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	config.AuthConfig
}

func NewAuthHanlder(authConfig config.AuthConfig) *AuthHandler {
	return &AuthHandler{AuthConfig: authConfig}
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
	uId, err := util.ParseToken(
		tokenString.Refresh_token,
		[]byte(handler.AuthConfig.JWTRefreshToken),
	)
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, err.Error())
	}
	accessToken := util.GenerateAccessClaims(
		uId,
		[]byte(handler.AuthConfig.JWTSecret),
		time.Duration(handler.AuthConfig.TokenExpire)*time.Minute,
	)
	newToken := new(entities.Auth)
	newToken.Token = accessToken
	newToken.RefreshToken = tokenString.Refresh_token
	return c.JSON(newToken)
}
