package main

import (
	"api-blog/api/config"
	"api-blog/api/gorm_repository"
	"api-blog/api/handler"
	"api-blog/api/middleware"
	"api-blog/api/routes"
	"api-blog/api/util"
	_ "api-blog/docs"
	"api-blog/pkg/entities"
	"api-blog/pkg/usecase"
	"api-blog/src/reaction"
	"fmt"
	"log"

	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var startTime = time.Now()

type HealthCheckResponse struct {
	Message   string  `json:"message"`
	Uptime    float64 `json:"uptime"`
	Timestamp int64   `json:"timestamp"`
}

// @title web Blog
// @version 1.0
// @description This is a web blog
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Apply "bearer " before token in authorization
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	//DB
	db, err := util.ConnectProstgrest(cfg)
	if err != nil {
		panic("Cannot connect Database: %v")
	}

	if err = db.
		AutoMigrate(
			&entities.User{},
			&entities.Post{},
			&entities.Slug{},
			&entities.Comment{},
			&entities.Reaction{},
		); err != nil {
		log.Panic("failed to migrate database: ", err)
	}

	//Minio
	minioClient, err := util.ConnectMinio(cfg)
	if err != nil {
		panic("Fail to load Minio")
	}

	//middlerware
	middle := middleware.NewJWTMiddleware(cfg.AuthConfig.JWTSecret)

	//register usecase
	authHandler := handler.NewAuthHanlder(cfg.AuthConfig.JWTSecret, cfg.AuthConfig.JWTRefreshToken)
	//user
	userRepo := gorm_repository.NewUserGormRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC, *cfg)

	//Media
	mediaHandler := handler.NewMediaHandler(*cfg, minioClient)

	// slug
	slugRepo := gorm_repository.NewSlugGormRepository(db)
	slugUC := usecase.NewSlugUseCase(slugRepo)

	//post
	postRepo := gorm_repository.NewPostGormRepository(db)
	postUC := usecase.NewPostUseCase(postRepo)
	postHandler := handler.NewPostHandler(postUC, slugUC, userUC)

	// comment
	commentRepo := gorm_repository.NewCommentGormRepository(db)
	commentUC := usecase.NewCommentUseCase(commentRepo)
	commentHandler := handler.NewCommentHandler(commentUC)

	//app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		c.Locals("userService", userUC)
		return c.Next()
	})
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Api Blog Metrics Page"}))
	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		elapsed_time := time.Since(startTime)
		uptime := elapsed_time.Seconds()
		res := HealthCheckResponse{Message: "Still alive lmao", Uptime: uptime, Timestamp: time.Now().UnixNano()}
		return c.JSON(res)
	})
	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	routes.UserRouter(v1, *userHandler, *middle)
	routes.AuthRouter(v1, *authHandler, *userHandler, *middle)
	routes.MediaRouter(v1, *mediaHandler, *middle)
	routes.PostRouter(v1, *postHandler, *middle)
	routes.CommentRouter(v1, *commentHandler, *middle)
	reaction.RegisterReactionApi(v1, *middle)

	port := fmt.Sprintf(":%v", cfg.Server.Port)
	app.Listen(port)
}
