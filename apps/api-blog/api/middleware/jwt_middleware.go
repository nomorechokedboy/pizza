package middleware

import (
	"api-blog/api/util"

	"github.com/gofiber/fiber/v2"
)

type JWTMiddleware struct {
	secret []byte
}

func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{secret: []byte(secret)}
}

func (m *JWTMiddleware) Protected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.ErrUnauthorized
		}
		tokenString := authHeader[len("Bearer "):]
		uId, err := util.ParseToken(tokenString, m.secret)
		if err != nil {
			return fiber.NewError(fiber.ErrUnauthorized.Code, "Invalid or missing Token")
		}
		c.Locals("uId", uId)
		return c.Next()

	}
}
