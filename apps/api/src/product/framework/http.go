package framework

import (
	"api/src/product/domain"
	"api/src/product/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

// InsertProduct godoc
// @Summary Create product api
// @Description Create a new product with coresponding body
// @Accept json
// @Produce json
// @Param todo body domain.ProductReq true "New Product"
// @Success 201 {object} domain.Product
// @Failure 400
// @Router /product/insert [post]
// @tags Product
func InsertProduct(ctx *fiber.Ctx) error {
	req := domain.ProductReq{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}

	useCase := ctx.Locals("insertProductUseCase").(*usecases.InsertProductUseCase)
	product, err := useCase.Execute(req)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(product)
}

// UpdateProduct godoc
// @Summary Update product api
// @Description Update a product with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "Product Id"
// @Param product body domain.ProductReq true "Update product"
// @Success 201 {object} domain.Product
// @Failure 400
// @Router /product/update/{id} [put]
// @tags Product
func Updateproduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	req := domain.ProductReq{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}

	useCase := ctx.Locals("updateProductUseCase").(*usecases.UpdateProductUseCase)
	product, err := useCase.Execute(uint(id), req)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(product)
}

// DeleteProduct godoc
// @Summary Delete product api
// @Description Delete a product with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "Product id"
// @Success 201 {object} domain.Product
// @Failure 400
// @Router /product/delete/{id} [delete]
// @tags Product
func DeleteProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	useCase := ctx.Locals("deleteProductUseCase").(*usecases.DeleteProductUseCase)
	product, err := useCase.Execute(uint(id))
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(product)
}

// FindProduct godoc
// @Summary Find categories api
// @Description Get a list of categories with coresponding query parameters
// @Accept json
// @Produce json
// @Param page query int false "Product page number"
// @Param pageSize query int false "Product page size return"
// @Param q query string false "Product query"
// @Param sort query string false "Sort direction" Enums(asc, desc) default(desc)
// @Param sortBy query string false "Sort by" Enums(id, name, description, sku, price, category_id) default(id)
// @Param categoryId query string false "Category ID"
// @Success 201 {object} common.BasePaginationResponse[domain.Product]
// @Failure 400
// @Router /product/find [get]
// @tags Product
func FindProduct(ctx *fiber.Ctx) error {
	queries := new(domain.ProductQuery)
	if err := ctx.QueryParser(queries); err != nil {
		return err
	}

	useCase := ctx.Locals("findProductUseCase").(*usecases.FindProductUseCase)
	products, err := useCase.Execute(queries)

	if err != nil {
		return err
	}

	return ctx.JSON(products)
}

// FindOneProduct godoc
// @Summary Find product details api
// @Description Get a product details with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "Product id"
// @Success 201 {object} domain.Product
// @Failure 400
// @Router /product/details/{id} [get]
// @tags Product
func FindOneProduct(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	useCase := ctx.Locals("findOneProductUseCase").(*usecases.FindOneProductUseCase)
	product, err := useCase.Execute(uint(id))
	if err != nil {
		return err
	}

	return ctx.JSON(product)
}
