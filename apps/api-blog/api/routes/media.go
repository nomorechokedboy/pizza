package routes

import (
	"api-blog/api/handler"

	"github.com/gofiber/fiber/v2"
)

func MediaRouter(app fiber.Router, handler handler.MediaHandler) {
	media := app.Group("/media")
	media.Post("/upload", handler.PostImage)
	media.Get("/:objectName", handler.GetMedia)
}
