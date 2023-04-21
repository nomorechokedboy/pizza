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

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
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
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic("Can't get db", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

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

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// We set the payload for the message.
	body := "Golang is awesome - Keep Moving Forward!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	// If there is an error publishing the message, a log will be displayed in the terminal.
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Congrats, sending message: %s", body)

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
	postHandler := handler.NewPostHandler(postUC, slugUC, userUC, cfg, minioClient)

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
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		c.Locals("userService", userUC)
		return c.Next()
	})
	prometheus := fiberprometheus.New("my-service-name")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
	app.Get("/healthCheck", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
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
