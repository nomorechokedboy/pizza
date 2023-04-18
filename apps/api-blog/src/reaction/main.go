package reaction

import (
	"api-blog/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterReactionApi(v1 fiber.Router, middlewares middleware.JWTMiddleware) {
	route := v1.Group("/reaction")

	route.Use(middlewares.IsAuth)
	route.Post("/react", ReactToEntity)
	route.Delete("/drop", DropEntityReaction)
}
