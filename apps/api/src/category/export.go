package category

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"api/src/category/infrastructure"
	"api/src/category/repository"

	"github.com/gofiber/fiber/v2"
)

var CategoryMemRepo = repository.CategoryInMemoryRepo{Data: make([]domain.Category, 0), IsErr: false}
var InsertCategoryUseCase = usecases.InsertCategoryUseCase{Repo: &CategoryMemRepo}
var UpdateCategoryUseCase = usecases.UpdateCategoryUseCase{Repo: &CategoryMemRepo}
var DeleteCategoryUseCase = usecases.DeleteCategoryUseCase{Repo: &CategoryMemRepo}
var FindCategoryUseCase = usecases.FindCategoryUseCase{Repo: &CategoryMemRepo}
var FindOneCategoryUseCase = usecases.FindOneCategoryUseCase{Repo: &CategoryMemRepo}

func New(v1 fiber.Router) {
	categoryRoute := v1.Group("/category")

	categoryRoute.Post("/insert", infrastructure.InsertCategory)
	categoryRoute.Put("/update/:id<int>", infrastructure.UpdateCategory)
	categoryRoute.Delete("/delete/:id<int>", infrastructure.DeleteCategory)
	categoryRoute.Get("/find", infrastructure.FindCategory)
	categoryRoute.Get("/details/:id<int>", infrastructure.FindOneCategory)
}
