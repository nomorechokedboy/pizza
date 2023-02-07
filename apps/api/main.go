package main

import (
	_ "api/docs"
	"log"
	"os"

	"api/src/category"
	CategoryInfrastructure "api/src/category/infrastructure"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func getEnv(name, fallback string) string {
	if env, ok := os.LookupEnv(name); ok {
		return env
	}

	return fallback
}

// @title Fiber Pizza API
// @version 1.0.0
// @description This is pizza api swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func main() {
	DB_DNS := getEnv("DB_DNS", "host=localhost user=postgres password=postgres dbname=pizza port=5432 sslmode=disable")
	db, err := gorm.Open(postgres.Open(DB_DNS), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err = db.AutoMigrate(&CategoryInfrastructure.Category{}); err != nil {
		panic("failed to migrate database")
	}

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		category.RegisterUseCases(c, db)
		return c.Next()
	})
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())

	category.New(v1)
	app.Get("healthz", HealthCheck)
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs")
	})

	log.Fatal(app.Listen(":3001"))
}
