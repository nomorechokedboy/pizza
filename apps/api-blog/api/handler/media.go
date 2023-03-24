package handler

import (
	"api-blog/api/config"
	"context"
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
// @Param id path string true "imageName"
// @Produce octet-stream
// @Success 200
// @Router /media/{objectName} [get]
func (handler *MediaHandler) GetMedia(c *fiber.Ctx) error {
	ctx := context.Background()
	objectName := c.Params("objectName")
	if objectName == "" {
		return fiber.NewError(fiber.ErrBadGateway.Code, "invalid object name")
	}
	newObject, err := handler.minioClient.GetObject(ctx, handler.config.Minio.BucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Can not find Image")
	}

	return c.SendStream(newObject)
}
