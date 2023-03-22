package framework

import (
	_ "api/src/cart/domain"
	CartItemDomain "api/src/cartItem/domain"
	"api/src/cartItem/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

// InsertCartItem godoc
// @Summary Create product api
// @Description Create a new product with coresponding body
// @Accept json
// @Produce json
// @Param uid path string true "User id"
// @Param todo body CartItemDomain.WriteCartItemRequest true "New Cart Item"
// @Success 201 {object} CartItemDomain.CartItem
// @Failure 400
// @Router /cart/insert/{uid} [post]
// @tags Cart
func InsertCartItem(ctx *fiber.Ctx) error {
	uid, err := ctx.ParamsInt("uid")
	if err != nil {
		return err
	}

	req := CartItemDomain.WriteCartItemRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}

	useCase := ctx.Locals("insertCartItemUseCase").(*usecases.InsertCartItemUseCase)
	cartItem, err := useCase.Execute(uint(uid), req)
	if err != nil {
		return ctx.Status(409).JSON(nil)
	}

	return ctx.Status(201).JSON(cartItem)
}

// FindUserCart godoc
// @Summary Find user's cart details api
// @Description Get user's cart details with coresponding user id
// @Accept json
// @Produce json
// @Param uid path string true "User id"
// @Success 201 {object} domain.Cart
// @Failure 400
// @Router /cart/details/{uid} [get]
// @tags Cart
func FindUserCart(ctx *fiber.Ctx) error {
	uid, err := ctx.ParamsInt("uid")
	if err != nil {
		return err
	}

	useCase := ctx.Locals("findCartUseCase").(*usecases.FindCartItemsUseCase)
	cart, err := useCase.Execute(uint(uid))
	if err != nil {
		return err
	}

	return ctx.JSON(cart)
}

// DeleteCartItem godoc
// @Summary Delete user's cart item api
// @Description Delete user's cart item with coresponding product id
// @Accept json
// @Produce json
// @Param uid path string true "User id"
// @Param productID path string true "Product id"
// @Success 201 {object} domain.Cart
// @Failure 400
// @Router /cart/delete/{uid}/{productID} [delete]
// @tags Cart
func DeleteCartItem(ctx *fiber.Ctx) error {
	uid, err := ctx.ParamsInt("uid")
	if err != nil {
		return err
	}

	productID, err := ctx.ParamsInt("productID")
	if err != nil {
		return err
	}

	useCase := ctx.Locals("deleteCartItemUseCase").(*usecases.DeleteCartItemUseCase)
	cart, err := useCase.Execute(uint(uid), uint(productID))
	if err != nil {
		return err
	}

	return ctx.JSON(cart)
}

// UpdateCartItem godoc
// @Summary Update user's cart item api
// @Description Update user's cart item with coresponding product id
// @Accept json
// @Produce json
// @Param uid path string true "User Id"
// @Param product body CartItemDomain.WriteCartItemRequest true "Update cart item"
// @Success 201 {object} domain.Cart
// @Failure 400
// @Router /cart/update/{uid} [put]
// @tags Cart
func UpdateCartItem(ctx *fiber.Ctx) error {
	uid, err := ctx.ParamsInt("uid")
	if err != nil {
		return err
	}

	req := CartItemDomain.WriteCartItemRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(500).JSON(nil)
	}

	useCase := ctx.Locals("updateCartItemUseCase").(*usecases.UpdateCartItemUseCase)
	cart, err := useCase.Execute(uint(uid), req)
	if err != nil {
		return err
	}

	return ctx.JSON(cart)
}
