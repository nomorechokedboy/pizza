package infrastructure

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

// InsertCategory godoc
// @Summary Create category api
// @Description Create a category with coresponding body
// @Accept json
// @Produce json
// @Param body body domain.WriteCategoryBody true "New Category"
// @Success 201 {object} domain.Category
// @Failure 400
// @Router /category/insert [post]
// @tags Category
func InsertCategory(ctx *fiber.Ctx) error {
	req := domain.WriteCategoryBody{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}

	useCase := ctx.Locals("insertCategoryUseCase").(*usecases.InsertCategoryUseCase)
	if useCase == nil {
		return ctx.Status(500).SendString("Internal error")
	}
	category, err := useCase.Execute(&req)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(category)
}

// UpdateCategory godoc
// @Summary Update category api
// @Description Update a category with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "Category Id"
// @Param category body domain.WriteCategoryBody true "Update Category"
// @Success 201 {object} domain.Category
// @Failure 400
// @Router /category/update/{id} [put]
// @tags Category
func UpdateCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	req := domain.WriteCategoryBody{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}

	useCase := ctx.Locals("updateCategoryUseCase").(*usecases.UpdateCategoryUseCase)
	category, err := useCase.Execute(&id, &req)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(category)
}

// DeleteCategory godoc
// @Summary Delete category api
// @Description Delete a category with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "Category id"
// @Success 201 {object} domain.Category
// @Failure 400
// @Router /category/delete/{id} [delete]
// @tags Category
func DeleteCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	useCase := ctx.Locals("deleteCategoryUseCase").(*usecases.DeleteCategoryUseCase)
	category, err := useCase.Execute(&id)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(category)
}

// FindCategory godoc
// @Summary Find categories api
// @Description Get a list of categories with coresponding query parameters
// @Accept json
// @Produce json
// @Param page query int false "Category page number"
// @Param pageSize query int false "Category page size return"
// @Param q query string false "Category query"
// @Success 201 {object} []domain.Category
// @Failure 400
// @Router /category/find [get]
// @tags Category
func FindCategory(ctx *fiber.Ctx) error {
	queries := new(domain.CategoryQuery)
	if err := ctx.QueryParser(queries); err != nil {
		return err
	}

	useCase := ctx.Locals("findCategoryUseCase").(*usecases.FindCategoryUseCase)
	categories, err := useCase.Execute(queries)

	if err != nil {
		return err
	}

	return ctx.JSON(categories)
}

// FindOneCategory godoc
// @Summary Find category details api
// @Description Get a category details with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "Category id"
// @Success 201 {object} domain.Category
// @Failure 400
// @Router /category/details/{id} [get]
// @tags Category
func FindOneCategory(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	useCase := ctx.Locals("findOneCategoryUseCase").(*usecases.FindOneCategoryUseCase)
	category, err := useCase.Execute(&id)
	if err != nil {
		return err
	}

	return ctx.JSON(category)
}
