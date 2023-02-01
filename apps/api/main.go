package main

import (
	_ "api/docs"
	"log"

	"api/src/inventory"
	"api/src/inventory/domain/usecases"
	"api/src/product"
	"api/src/product/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var createProductUseCase = domain.CreateProductUseCase{Repo: &product.ProductMemRepo}
var DeleteInventoryUseCase = usecases.DeleteInventoryUseCase{Repo: &inventory.InventoryMemRepo}

// HealthCheck godoc
// @Summary healthcheck api
// @Description Get uptime and application status
// @Accept json
// @Produce json
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// @title Fiber Pizza API
// @version 1.0.0
// @description This is pizza api swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3001
// @BasePath /
func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("createProductUseCase", createProductUseCase)
		return c.Next()
	})
	app.Post("/create-product", product.CreateProduct)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", HealthCheck)

	log.Fatal(app.Listen(":3001"))
}
