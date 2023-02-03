package user

import (
	"api/src/user/domain"
	"api/src/user/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc

func CreateUser(ctx *fiber.Ctx) error {
	req := domain.CreateUserReq{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}

	useCase := ctx.Locals("CreateUserUseCase").(usecases.CreateUserUseCase)
	user, err := useCase.Execute(&req)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("Id")
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}

	req := domain.CreateUserReq{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}
	useCase := ctx.Locals("updateUserUseCase").(usecases.UpdateUserUseCase)
	user, err := useCase.Execute(&id, &req)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(user)
}
