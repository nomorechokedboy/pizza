package routes

import (
	"api-blog/api/handler"

	"github.com/gofiber/fiber/v2"
)

func TokenRouter(app fiber.Router, handler handler.AuthHandler) {
	token := app.Group("token")
	token.Get("/refresh_token", handler.RefreshToken)
}
