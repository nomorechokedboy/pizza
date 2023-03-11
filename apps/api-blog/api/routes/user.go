package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, handler handler.UserHandler, jwtSecret string, jwtRefreshSecret string) {
	users := app.Group("/users")
	publicRouter(users, handler)
	privateRouter(users, handler, jwtSecret)
}

func publicRouter(app fiber.Router, handler handler.UserHandler) {
	app.Post("/register", handler.CreateUser)
	app.Post("/login", handler.Login)
	app.Post("/forgot-password", handler.ForgotPassword)

}

func privateRouter(app fiber.Router, handler handler.UserHandler, jwtSecret string) {
	middle := middleware.NewJWTMiddleware(jwtSecret)
	app.Use(middle.Protected())
	app.Get("/", handler.GetAuthUserById)
	app.Put("/update", handler.UpdateUserById)
	app.Put("/reset-password", handler.ResetPassword)
}
