package framework

import (
	"api/src/product/domain"
	"api/src/product/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

// CreateProduct godoc
// @Summary Create product api
// @Description Create a new product with coresponding body
// @Accept json
// @Produce json
// @Param todo body domain.ProductReq true "New Product"
// @Success 201 {object} domain.Product
// @Failure 400
// @Router /create-product [post]
func CreateProduct(ctx *fiber.Ctx) error {
	req := domain.ProductReq{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}

	useCase := ctx.Locals("createProductUseCase").(usecases.CreateProductUseCase)
	product, err := useCase.Execute(&req)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(product)
}
