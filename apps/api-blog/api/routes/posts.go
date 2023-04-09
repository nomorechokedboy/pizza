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
	app.Get("/", handler.GetAllPosts)
	app.Get("/:slug<min(1)>", handler.GetPostBySlug)
}

func privatePostRouter(app fiber.Router, handler handler.PostHandler, middle middleware.JWTMiddleware) {
	app.Use(middle.Protected())
	app.Post("/", handler.CreatePost)
	app.Put("/:id", handler.UpdatePost)
	app.Delete("/:id", handler.DeletePost)
}
