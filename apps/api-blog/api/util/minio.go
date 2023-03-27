package util

import (
	"api-blog/api/config"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectMinio(cfg *config.Config) (*minio.Client, error) {
	minioClient, err := minio.New(cfg.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connect Minio success")
	return minioClient, nil
}
