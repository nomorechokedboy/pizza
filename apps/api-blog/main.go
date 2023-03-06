package main

import (
	"api-blog/api/config"
	"api-blog/api/gorm_repository"
	"api-blog/api/handler"
	"api-blog/api/routes"
	"api-blog/api/util"
	_ "api-blog/docs"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"fmt"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	//DB
	db, err := util.ConnectProstgrest(cfg)
	if err != nil {
		log.Fatalf("Cannot connect Database: %v", err)
	}
	db.AutoMigrate(&entities.User{})

	//register usecase
	//user
	userRepo := gorm_repository.NewUserGormRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC, cfg.AuthConfig.JWTSecret, cfg.AuthConfig.JWTRefreshToken)

	//app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
	app.Get("/swagger/doc.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	routes.UserRouter(v1, *userHandler, cfg.AuthConfig.JWTSecret, cfg.AuthConfig.JWTRefreshToken)
	port := fmt.Sprintf(":%v", cfg.Server.Port)
	app.Listen(port)
	log.Printf("Server started on port %v", cfg.Server.Port)

}
