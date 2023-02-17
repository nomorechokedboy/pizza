package infrastructure

import (
	"api/src/user/domain"
	"api/src/user/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc
// @Summary Create user api
// @Description Create a user with coresponding body
// @Accept json
// @Produce json
// @Param body body domain.CreateUserReq true "New User"
// @Success 201 {object} domain.User
// @Failure 400
// @Router /user/create [post]
// @tags User
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

// UpdateUser godoc
// @Summary Update user api
// @Description Update a user with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Param user body domain.CreateUserReq true "Update User"
// @Success 201 {object} domain.User
// @Failure 400
// @Router /user/update/{id} [put]
// @tags user
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

// DeleteUser godoc
// @Summary Delete user api
// @Description Delete a user with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 201 {object} domain.User
// @Failure 400
// @Router /user/delete/{id} [delete]
// @tags User
func DeleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	useCase := ctx.Locals("deleteUserUseCase").(usecases.DeleteUserUseCase)
	user, err := useCase.Execute(&id)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(user)
}

// FindUser godoc
// @Summary Find users api
// @Description Get a list of users with coresponding query parameters
// @Accept json
// @Produce json
// @Param page query int false "User page number"
// @Param pageSize query int false "User  size return"
// @Param q query string false "User query"
// @Success 201 {object} []domain.User
// @Failure 400
// @Router /user/find [get]
// @tags user
func FindUser(ctx *fiber.Ctx) error {
	queries := new(domain.UserQuery)
	if err := ctx.QueryParser(queries); err != nil {
		return err
	}

	useCase := ctx.Locals("findUserUseCase").(*usecases.FindUserUseCase)
	user, err := useCase.Execute(queries)

	if err != nil {
		return err
	}

	return ctx.JSON(user)
}

// FindOneUser godoc
// @Summary Find user details api
// @Description Get a user details with coresponding id
// @Accept json
// @Produce json
// @Param id path string true "User id"
// @Success 201 {object} domain.User
// @Failure 400
// @Router /user/details/{id} [get]
// @tags User
func FindOneUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	useCase := ctx.Locals("findOneUserUseCase").(*usecases.FindOneUserUseCase)
	user, err := useCase.Execute(&id)
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}
