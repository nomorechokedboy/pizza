package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func MediaRouter(app fiber.Router, handler handler.MediaHandler, middleware middleware.JWTMiddleware) {
	media := app.Group("/media")
	media.Get("/:uuId/:objectName", handler.GetMedia)
	media.Use(middleware.Protected())
	media.Post("/upload", handler.PostImage)
}
