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
	"fmt"
	"log"

	// "context"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	// "github.com/minio/minio-go/v7"
	// "github.com/minio/minio-go/v7/pkg/credentials"
)

// @title web Blog
// @version 1.0
// @description This is a web blog
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Apply "bearer " before token in authorization
func main() {

	//
	// TODO: Refactor this minio connection code
	// ctx := context.Background()
	// endpoint := "localhost:9000"
	// accessKeyID := "admin"
	// secretAccessKey := "admin123"
	// useSSL := false

	// // Initialize minio client object.
	// minioClient, err := minio.New(endpoint, &minio.Options{
	// 	Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	// 	Secure: useSSL,
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // Change this code since it won't work in your local machine
	// bucketName := "general"
	// objectName := "todo-actix.zip"
	// filePath := "./tmp/todo-actix.zip"
	// contentType := "application/zip"

	// // Upload the zip file with FPutObject
	// info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Sze)
	//
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
	db.AutoMigrate(&entities.Post{})
	db.AutoMigrate(&entities.Slug{})

	//middlerware

	middle := middleware.NewJWTMiddleware(cfg.AuthConfig.JWTSecret)

	//register usecase
	authHandler := handler.NewAuthHanlder(cfg.AuthConfig.JWTSecret, cfg.AuthConfig.JWTRefreshToken)
	//user
	userRepo := gorm_repository.NewUserGormRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC, *cfg)

	// slug
	slugRepo := gorm_repository.NewSlugGormRepository(db)
	slugUC := usecase.NewSlugUseCase(slugRepo)

	//post
	postRepo := gorm_repository.NewPostGormRepository(db)
	postUC := usecase.NewPostUseCase(postRepo)
	postHandler := handler.NewPostHandler(postUC, slugUC)

	//app
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("Helo, world")
	})

	app.Get("/docs/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	routes.UserRouter(v1, *userHandler, *middle)
	routes.AuthRouter(v1, *authHandler, *userHandler, *middle)
	routes.PostRouter(v1, *postHandler, *middle)

	port := fmt.Sprintf(":%v", cfg.Server.Port)
	log.Printf("Server started on port %v", cfg.Server.Port)
	app.Listen(port)
}
