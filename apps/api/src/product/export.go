package product

import (
	"api/src/common"
	"api/src/product/domain"
	"api/src/product/domain/usecases"
	"api/src/product/framework"
	"api/src/product/repository"
	ProductRepo "api/src/product/repository/gorm"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var ProductMemRepo = repository.ProductInMemoryRepo{DataStore: make([]domain.Product, 0), IsErr: false}

func New(v1 fiber.Router) {
	productRoute := v1.Group("/product")

	productRoute.Post("/insert", framework.InsertProduct)
	productRoute.Put("/update/:id<int;min(0)>", framework.Updateproduct)
	productRoute.Delete("/delete/:id<int;min(0)>", framework.DeleteProduct)
	productRoute.Get("/details/:id<int;min(0)>", framework.FindOneProduct)
	productRoute.Get("/find", framework.FindProduct)
}

func RegisterUseCases(c *fiber.Ctx, db *gorm.DB) {
	Repo := &ProductRepo.ProductGormRepo{Conn: db}
	insertProductUseCase := usecases.InsertProductUseCase{Repo: Repo, Validator: &common.ValidatorAdapter}
	updateProductUseCase := usecases.UpdateProductUseCase{Repo: Repo}
	deleteProductUseCase := usecases.DeleteProductUseCase{Repo: Repo}
	findOneProductUseCase := usecases.FindOneProductUseCase{Repo: Repo}
	findProductUseCase := usecases.FindProductUseCase{Repo: Repo}

	c.Locals("insertProductUseCase", &insertProductUseCase)
	c.Locals("updateProductUseCase", &updateProductUseCase)
	c.Locals("deleteProductUseCase", &deleteProductUseCase)
	c.Locals("findProductUseCase", &findProductUseCase)
	c.Locals("findOneProductUseCase", &findOneProductUseCase)
}
