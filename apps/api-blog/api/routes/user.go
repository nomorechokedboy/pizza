package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"
	"api-blog/api/util"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, handler handler.UserHandler) {
	users := app.Group("/users")
	users.Post("/register", handler.CreateUser)
	users.Get("/refresh_access_token", util.RefreshAccessToken)
	users.Post("/login", handler.Login)
	middle := middleware.NewJWTMiddleware("my-secret")
	users.Use(middle.Protected())
	users.Get("/userid", handler.GetUserById)
	users.Post("/Updated", handler.UpdateUserById)
}
