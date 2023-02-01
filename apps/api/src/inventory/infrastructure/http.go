package infrastructure

import (
	"api/src/inventory/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

// DeleteProduct godoc
// @Summary Delete product api
// @Description Delete a product with coresponding id
// @Accept json
// @Produce json
// @Param id path string false "Inventory id"
// @Success 201 {object} domain.Inventory
// @Failure 400
// @Router /delete-inventory [delete]
func DeleteProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	useCase := ctx.Locals("deleteProductUseCase").(usecases.DeleteInventoryUseCase)
	product, err := useCase.Execute(&id)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(product)
}
