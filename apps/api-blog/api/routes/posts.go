package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func PostRouter(app fiber.Router, handler handler.PostHandler, jwtSecret middleware.JWTMiddleware) {
	users := app.Group("/posts")
	publicPostRouter(users, handler)
	privatePostRouter(users, handler, jwtSecret)

}

func publicPostRouter(app fiber.Router, handler handler.PostHandler) {
	app.Get("/", handler.GetAllPostsByUserID)
	app.Get("/:slug", handler.GetPostBySlug)
}

func privatePostRouter(app fiber.Router, handler handler.PostHandler, jwtSecret middleware.JWTMiddleware) {
	app.Post("/", handler.CreatePost)
	app.Put("/", handler.UpdatePost)
	app.Delete("/", handler.DeletePost)
}
