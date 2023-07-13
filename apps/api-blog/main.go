package main

import (
	"api-blog/api/config"
	"api-blog/api/util"
	_ "api-blog/docs"
	"api-blog/pkg/entities"
	notificationEntities "api-blog/src/notification/entities"
	"api-blog/src/server"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
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

	app := server.New(cfg, db, minioClient, rdb)
	port := fmt.Sprintf(":%v", cfg.Server.Port)
	app.Listen(port)
}
