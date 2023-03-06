package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, handler handler.UserHandler, jwtSecret string, jwtRefreshSecret string) {
	users := app.Group("/users")
	publicRouter(users, handler, jwtRefreshSecret)
	privateRouter(users, handler, jwtSecret)

}

func publicRouter(app fiber.Router, handler handler.UserHandler, jwtRefreshSecret string) {
	app.Post("/register", handler.CreateUser)
	app.Get("/refresh_access_token", handler.RefreshAccessToken)
	app.Post("/login", handler.Login)
}

func privateRouter(app fiber.Router, handler handler.UserHandler, jwtSecret string) {
	middle := middleware.NewJWTMiddleware(jwtSecret)
	app.Use(middle.Protected())
	app.Get("/", handler.GetAuthUserById)
	app.Post("/update", handler.UpdateUserById)
}
