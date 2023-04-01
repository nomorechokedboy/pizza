package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database struct {
		Port     string `envconfig:"DB_PORT" default:"5432"`
		Host     string `envconfig:"DB_HOST" default:"localhost"`
		Name     string `envconfig:"DB_NAME" default:"pizza"`
		User     string `envconfig:"DB_USER" default:"postgres"`
		Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	}
	Server struct {
		Host string `envconfig:"HOST" default:""`
		Port string `envconfig:"PORT" default:"8080"`
	}
	AuthConfig struct {
		JWTRefreshToken string `envconfig:"JWT_REFRESH_SECRET" default:"refresh-secret"`
		JWTSecret       string `envconfig:"JWT_SECRET" default:"my-secret"`
	}
	AuthEmail struct {
		Email    string `envconfig:"EMAIL" default:"kristiannguyen276@gmail.com"`
		Password string `envconfig:"EMAIL_PASSWORD" default:"figjbdfsggwhcvbr"`
	}
	AppAPI struct {
		Link string `envconfig:"API_LINK" default:"Hello"`
	}
	Minio struct {
		EndPoint        string `envconfig:"END_POINT" default:"localhost:9000"`
		AccessKeyID     string `envconfig:"ACCESSKEYID" default:"admin"`
		SecretAccessKey string `envconfig:"SECRET_ACCESS_KEY" default:"admin123"`
		UseSSL          bool   `envconfig:"USESSL" default:"false"`
		BucketName      string `envconfig:"BUCKET_NAME" default:"general"`
	}
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
