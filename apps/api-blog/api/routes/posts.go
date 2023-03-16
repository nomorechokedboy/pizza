package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func PostRouter(app fiber.Router, handler handler.PostHandler, jwtSecret string) {
	users := app.Group("/posts")
	publicPostRouter(users, handler)
	privatePostRouter(users, handler, jwtSecret)

}

func publicPostRouter(app fiber.Router, handler handler.PostHandler) {
	app.Get("/", handler.GetAllPosts)
	app.Get("/", handler.GetAllPostsByUserID)
	app.Get("/:slug", handler.GetPostBySlug)
}

func privatePostRouter(app fiber.Router, handler handler.PostHandler, jwtSecret string) {
	middle := middleware.NewJWTMiddleware(jwtSecret)
	app.Use(middle.Protected())
	app.Post("/", handler.CreatePost)
	app.Put("/", handler.UpdatePost)
	app.Delete("/", handler.DeletePost)
}
