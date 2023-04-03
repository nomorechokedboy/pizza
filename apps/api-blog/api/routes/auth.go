package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, authHandler handler.AuthHandler, userHandler handler.UserHandler, middlerware middleware.JWTMiddleware) {
	auth := app.Group("/auth")
	publicAuthRouter(auth, userHandler, authHandler)
	privateAuthRouter(auth, userHandler, middlerware)
}

func publicAuthRouter(app fiber.Router, userHandler handler.UserHandler, authHandler handler.AuthHandler) {
	app.Post("/login", userHandler.Login)
	app.Put("/reset-password", userHandler.ResetPassword)
	app.Post("/refresh-token", authHandler.RefreshToken)
	app.Post("/forgot-password", userHandler.ForgotPassword)
	app.Post("/register", userHandler.CreateUser)
}

func privateAuthRouter(app fiber.Router, userHandler handler.UserHandler, middleware middleware.JWTMiddleware) {
	app.Use(middleware.Protected())
	app.Get("/me", userHandler.GetAuthUserById)
	app.Put("/update-password", userHandler.UpdatePassword)
}
