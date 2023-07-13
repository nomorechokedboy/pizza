package handler

import (
	"api-blog/api/config"
	"api-blog/api/util"
	"api-blog/pkg/entities"
	notificationEntities "api-blog/src/notification/entities"
	"api-blog/src/server"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// DB
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
			&notificationEntities.NotificationObject{},
			&notificationEntities.Notification{},
			&notificationEntities.NotificationChange{},
		); err != nil {
		log.Panic("failed to migrate database: ", err)
	}
	// Minio
	minioClient, err := util.ConnectMinio(cfg)
	if err != nil {
		panic("Fail to load Minio")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.URI,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	handler(cfg, db, minioClient, rdb).ServeHTTP(w, r)
}

func handler(
	config *config.Config,
	db *gorm.DB,
	minioClient *minio.Client,
	rdb *redis.Client,
) http.HandlerFunc {
	app := server.New(config, db, minioClient, rdb)

	return adaptor.FiberApp(app)
}
