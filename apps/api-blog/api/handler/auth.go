package handler

import (
	"time"

	"api-blog/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
// @Tags Token
// @Produce json
// @Success 200 {object} string
// @Security ApiKeyAuth
// @Router /api/v1/token/refresh_token [get]
func (handler *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(handler.JwtRefreshToken), nil
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
	accessToken := util.GenerateAccessClaims(uId, []byte(handler.JwtSecret), 15*time.Minute)
	return c.JSON(accessToken)
}
