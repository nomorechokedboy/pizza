package routes

import (
	"api-blog/api/handler"
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func CommentRouter(app fiber.Router, handler handler.CommentHandler, jwtSecret middleware.JWTMiddleware) {
	users := app.Group("/comments")
	publicCommentRouter(users, handler)
	privateCommentRouter(users, handler, jwtSecret)

}

func publicCommentRouter(app fiber.Router, handler handler.CommentHandler) {
	app.Get("/", handler.GetAllComments)
}

func privateCommentRouter(app fiber.Router, handler handler.CommentHandler, middle middleware.JWTMiddleware) {
	app.Use(middle.Protected())
	app.Post("/", handler.CreateComment)
	app.Put("/:id", handler.UpdateComment)
	app.Delete("/:id", handler.DeleteComment)
}
