package cartitem

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"

	"api/src/cart/framework"
	cartRepo "api/src/cart/repository/redis"
	"api/src/cartItem/domain/usecases"
	"api/src/common"
)

func New(v1 fiber.Router) {
	cartRoute := v1.Group("/cart")

	cartRoute.Post("/insert/:uid<int;min(0)>", framework.InsertCartItem)
	cartRoute.Get("/details/:uid<int;min(0)>", framework.FindUserCart)
	cartRoute.Delete("/delete/:uid<int;min(0)>/:productID<int;min(0)>", framework.DeleteCartItem)
	cartRoute.Put("/update/:uid<int;min(0)>", framework.UpdateCartItem)
}

func RegisterUseCases(c *fiber.Ctx, rdb *redis.Client) {
	Repo := &cartRepo.CartRedisRepo{Conn: rdb}
	insertCartUseCase := usecases.InsertCartItemUseCase{Repo: Repo, Validator: &common.ValidatorAdapter}
	findCartUseCase := usecases.FindCartItemsUseCase{Repo: Repo}
	deleteCartItemUseCase := usecases.DeleteCartItemUseCase{Repo: Repo}
	updateCartItemUseCase := usecases.UpdateCartItemUseCase{Repo: Repo}

	c.Locals("insertCartItemUseCase", &insertCartUseCase)
	c.Locals("findCartUseCase", &findCartUseCase)
	c.Locals("deleteCartItemUseCase", &deleteCartItemUseCase)
	c.Locals("updateCartItemUseCase", &updateCartItemUseCase)
}
