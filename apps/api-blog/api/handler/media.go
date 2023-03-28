package handler

import (
	"api-blog/api/config"
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
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
	objectName := strings.ReplaceAll(file.Filename, " ", "")
	fileBuffer := buffer
	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	_, minioError := handler.minioClient.StatObject(ctx, handler.config.Minio.BucketName, objectName, minio.StatObjectOptions{})
	if minioError == nil {
		index := 0
		splitImage := strings.Split(objectName, ".")
		var newName string
		flag := true
		for flag {
			index += 1
			newName = fmt.Sprintf("%s-%d.%s", splitImage[0], index, splitImage[1])
			_, err := handler.minioClient.StatObject(ctx, handler.config.Minio.BucketName, newName, minio.StatObjectOptions{})
			if err != nil {
				flag = false
			}
		}
		objectName = newName
	}
	_, err = handler.minioClient.PutObject(ctx, handler.config.Minio.BucketName, objectName, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(objectName)
}

// @GetMedia godoc
// @Summary get Media
// @Tags Media
// @Accept json
// @Param objectName path string true "imageName"
// @Produce png
// @Success 200
// @Failure 404  "Cannot found the Image"
// @Failure 500 "Cannot get Image"
// @Router /media/{objectName} [get]
func (handler *MediaHandler) GetMedia(c *fiber.Ctx) error {
	ctx := context.Background()
	objectName := c.Params("objectName")
	if objectName == "" {
		return fiber.NewError(fiber.ErrBadGateway.Code, "invalid object name")
	}
	_, err := handler.minioClient.StatObject(ctx, handler.config.Minio.BucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Can not found Image")
	}
	newObject, err := handler.minioClient.GetObject(ctx, handler.config.Minio.BucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Can not get Image")
	}
	c.Set("Content-Type", "image/png")
	return c.SendStream(newObject)
}
