package routes

import (
	"api-blog/api/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, handler handler.UserHandler) {
	users := app.Group("/users")
	users.Post("/register", handler.CreateUser)
	users.Post("/login", handler.Login)
}
