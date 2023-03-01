package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// TODO: Refactor this minio connection code
	ctx := context.Background()
	endpoint := "localhost:9000"
	accessKeyID := "admin"
	secretAccessKey := "admin123"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Change this code since it won't work in your local machine
	bucketName := "general"
	objectName := "todo-actix.zip"
	filePath := "./tmp/todo-actix.zip"
	contentType := "application/zip"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	fmt.Println("Hello Cho vien")
}
