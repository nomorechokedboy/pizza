package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, handler handler.UserHandler, middleware middleware.JWTMiddleware) {
	users := app.Group("/user")
	publicRouter(users, handler)
	privateRouter(users, handler, middleware)
}

func publicRouter(app fiber.Router, handler handler.UserHandler) {
	app.Get("/:id", handler.FindUserById)
}

func privateRouter(app fiber.Router, handler handler.UserHandler, middle middleware.JWTMiddleware) {
	app.Use(middle.Protected())
	app.Put("/update", handler.UpdateUserById)
}
