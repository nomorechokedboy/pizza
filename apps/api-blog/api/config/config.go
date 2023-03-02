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
		Host string `env:"HOST" env-default:""`
		Port string `env:"PORT" env-default:"3001"`
	}
	AuthConfig struct {
		JWTSecret     string `envconfig:"JWT_SECRET" default:"my-secret"`
		JWTExpiration int64  `envconfig:"JWT_EXPIRATION" default:"3600"`
	}
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
