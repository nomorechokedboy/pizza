package user

import (
	"api/src/user/domain"
	"api/src/user/domain/usecases"
	"api/src/user/infrastructure"
	"api/src/user/repository"
	GormRepo "api/src/user/repository/gorm"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var UserMemRepo = repository.UserInMemoryRepo{UserList: make([]domain.User, 0), IsErr: false}
var CreateUserUseCase = usecases.CreateUserUseCase{Repo: &UserMemRepo}
var UpdateUserUseCase = usecases.FindUserUseCase{Repo: &UserMemRepo}
var DeleteUserUseCase = usecases.DeleteUserUseCase{Repo: &UserMemRepo}
var FindUserUseCase = usecases.FindUserUseCase{Repo: &UserMemRepo}
var FindOneUserUseCase = usecases.FindOneUserUseCase{Repo: &UserMemRepo}

func New(v1 fiber.Router) {
	userRoute := v1.Group("/user")

	userRoute.Post("/create", infrastructure.CreateUser)
	userRoute.Put("/update/:id<int>", infrastructure.UpdateUser)
	userRoute.Delete("/delete/:id<int>", infrastructure.DeleteUser)
	userRoute.Get("/find", infrastructure.FindUser)
	userRoute.Get("/details/:id<int>", infrastructure.FindOneUser)

}

func RegisterUseCases(c *fiber.Ctx, db *gorm.DB) {
	Repo := &GormRepo.UserGormRepo{DB: db}
	createUserUseCase := usecases.CreateUserUseCase{Repo: Repo}
	updateUserUseCase := usecases.UpdateUserUseCase{Repo: Repo}
	deleteUserUseCase := usecases.DeleteUserUseCase{Repo: Repo}
	findUserUseCase := usecases.FindUserUseCase{Repo: Repo}
	findOneUserUseCase := usecases.FindOneUserUseCase{Repo: Repo}

	c.Locals("createUserUseCase", &createUserUseCase)
	c.Locals("updateUserUseCase", &updateUserUseCase)
	c.Locals("deleteUserUseCase", &deleteUserUseCase)
	c.Locals("findUserUseCase", &findUserUseCase)
	c.Locals("findOneUserUseCase", &findOneUserUseCase)

}
