package handler

import (
	"api-blog/api/config"
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type MediaHandler struct {
	config      config.Config
	minioClient minio.Client
}

func NewMediaHandler(config config.Config, mionioClient *minio.Client) *MediaHandler {
	return &MediaHandler{
		config:      config,
		minioClient: *mionioClient,
	}
}

// @PostImage godoc
// @Summary Post to save image
// @Tags Media
// @ID	image
// @Produce		json
// @Accept	multipart/form-data
// @Security ApiKeyAuth
// @Param image formData file true "upfile"
// @Success 200
// @Router /media/upload [post]
func (handler *MediaHandler) PostImage(c *fiber.Ctx) error {
	ctx := context.Background()
	file, err := c.FormFile("image")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	buffer, err := file.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	fileName := strings.ReplaceAll(file.Filename, " ", "")
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	bucket := uuid.NewString()
	createBucketErr := handler.minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
	if createBucketErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "something bad happended")
	}
	_, err = handler.minioClient.PutObject(ctx, bucket, fileName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	link := fmt.Sprint(bucket + "/" + fileName)
	return c.JSON(link)
}

// @GetMedia godoc
// @Summary get Media
// @Tags Media
// @Accept json
// @Param uuId path string true "ID"
// @Param objectName path string true "object name"
// @Produce png
// @Success 200
// @Failure 404  "Cannot found the Image"
// @Failure 500 "Cannot get Image"
// @Router /media/{uuId}/{objectName} [get]
func (handler *MediaHandler) GetMedia(c *fiber.Ctx) error {
	ctx := context.Background()
	objectName := c.Params("objectName")
	bucket := c.Params("uuId")
	if objectName == "" {
		return fiber.NewError(fiber.ErrBadGateway.Code, "invalid object name")
	}
	_, err := handler.minioClient.StatObject(ctx, bucket, objectName, minio.StatObjectOptions{})
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Can not found Image")
	}
	newObject, err := handler.minioClient.GetObject(ctx, bucket, objectName, minio.GetObjectOptions{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Can not get Image")
	}
	c.Set("Content-Type", "image/png")
	return c.SendStream(newObject)
}
