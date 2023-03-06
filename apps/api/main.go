package main

import (
	_ "api/docs"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	Cart "api/src/cart"
	"api/src/category"
	CategoryDomain "api/src/category/domain"
	"api/src/config"
	InventoryDomain "api/src/inventory/domain"
	"api/src/product"
	ProductDomain "api/src/product/domain"
	"api/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
// @BasePath /api/v1
func main() {
	config, e := config.LoadEnv()
	if e != nil {
		log.Fatal(e)
	}
	DB_DNS := utils.GetDbURI(config)
	serverAddr := utils.GetServerAddress(config)
	db, err := gorm.Open(postgres.Open(DB_DNS), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	if err = db.AutoMigrate(&CategoryDomain.Category{}, &InventoryDomain.Inventory{}, &ProductDomain.Product{}); err != nil {
		panic("failed to migrate database")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		Cart.RegisterUseCases(c, rdb)
		category.RegisterUseCases(c, db)
		product.RegisterUseCases(c, db)
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

	Cart.New(v1)
	category.New(v1)
	product.New(v1)
	app.Get("/healthz", HealthCheck)
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs")
	})

	go func() {
		if appErr := app.Listen(serverAddr); appErr != nil {
			log.Panic(appErr)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
