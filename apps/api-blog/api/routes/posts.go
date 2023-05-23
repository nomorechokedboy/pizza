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
	app.Get("/:slug", handler.GetPostBySlug)
	app.Get("/t2s/:postSlug<minLen(1)>", handler.GetPostAudio)
}

func privatePostRouter(
	app fiber.Router,
	handler handler.PostHandler,
	middle middleware.JWTMiddleware,
) {
	app.Use(middle.IsAuth)
	app.Post("/", handler.CreatePost)
	app.Put("/:id", handler.UpdatePost)
	app.Delete("/:id", handler.DeletePost)
}
