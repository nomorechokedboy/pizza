package category

import (
	"api/src/category/domain"
	"api/src/category/domain/usecases"
	"api/src/category/infrastructure"
	"api/src/category/repository"
	GormRepo "api/src/category/repository/gorm"
	"api/src/common"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var CategoryMemRepo = repository.CategoryInMemoryRepo{Data: make([]domain.Category, 0), IsErr: false}
var InsertCategoryUseCase = usecases.InsertCategoryUseCase{Repo: &CategoryMemRepo, Validator: &common.ValidatorAdapter}
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

func RegisterUseCases(c *fiber.Ctx, db *gorm.DB) {
	Repo := &GormRepo.CategoryGormRepo{DB: db}
	insertCategoryUseCase := usecases.InsertCategoryUseCase{Repo: Repo, Validator: &common.ValidatorAdapter}
	updateCategoryUseCase := usecases.UpdateCategoryUseCase{Repo: Repo}
	deleteCategoryUseCase := usecases.DeleteCategoryUseCase{Repo: Repo}
	findOneCategoryUseCase := usecases.FindOneCategoryUseCase{Repo: Repo}
	findCategoryUseCase := usecases.FindCategoryUseCase{Repo: Repo}

	c.Locals("insertCategoryUseCase", &insertCategoryUseCase)
	c.Locals("updateCategoryUseCase", &updateCategoryUseCase)
	c.Locals("deleteCategoryUseCase", &deleteCategoryUseCase)
	c.Locals("findCategoryUseCase", &findCategoryUseCase)
	c.Locals("findOneCategoryUseCase", &findOneCategoryUseCase)
}
