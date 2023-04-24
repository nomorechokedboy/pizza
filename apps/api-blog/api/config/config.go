package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

/* type Config struct {
	Env      string `mapstructure:"ENV"`
	Database struct {
		Port     string `mapstructure:"DB_PORT"`
		Host     string `mapstructure:"DB_HOST"`
		Name     string `mapstructure:"DB_NAME"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
	}
	Server struct {
		Host string `mapstructure:"HOST"`
		Port string `mapstructure:"PORT"`
	}
	AuthConfig struct {
		JWTRefreshToken string `mapstructure:"JWT_REFRESH_SECRET"`
		JWTSecret       string `mapstructure:"JWT_SECRET"`
	}
	AuthEmail struct {
		Email    string `mapstructure:"EMAIL"`
		Password string `mapstructure:"EMAIL_PASSWORD"`
	}
	AppAPI struct {
		Link string `mapstructure:"API_LINK"`
	}
	AudioAPI struct {
		Link string `mapstructure:"LINK"`
		Key  string `mapstructure:"KEY"`
	}
	Minio struct {
		EndPoint        string `mapstructure:"END_POINT"`
		AccessKeyID     string `mapstructure:"ACCESSKEYID" `
		SecretAccessKey string `mapstructure:"SECRET_ACCESS_KEY"`
		UseSSL          bool   `mapstructure:"USESSL"`
		BucketName      string `mapstructure:"BUCKET_NAME"`
	}
} */

type Config struct {
	Env      string `env:"ENV" env-default:"dev"`
	Database struct {
		Port     string `env:"DB_PORT" env-default:"5432"`
		Host     string `env:"DB_HOST" env-default:"localhost"`
		Name     string `env:"DB_NAME" env-default:"pizza"`
		User     string `env:"DB_USER" env-default:"postgres"`
		Password string `env:"DB_PASSWORD" env-default:"postgres"`
	}
	Server struct {
		Host string `env:"HOST" env-default:""`
		Port string `env:"PORT" env-default:"8080"`
	}
	AuthConfig struct {
		JWTRefreshToken string `env:"JWT_REFRESH_SECRET" env-default:"refresh-secret"`
		JWTSecret       string `env:"JWT_SECRET" env-default:"token-secret"`
	}
	AuthEmail struct {
		Email    string `env:"EMAIL"`
		Password string `env:"EMAIL_PASSWORD"`
	}
	AppAPI struct {
		Link string `env:"FE_URL" env-default:"pizza-web-nuxt.vercel.app"`
	}
	AudioAPI struct {
		Link string `env:"LINK"`
		Key  string `env:"KEY"`
	}
	Minio struct {
		EndPoint        string `env:"END_POINT" env-default:"localhost:9000"`
		AccessKeyID     string `env:"ACCESSKEYID" env-default:"admin"`
		SecretAccessKey string `env:"SECRET_ACCESS_KEY" env-default:"admin123"`
		UseSSL          bool   `env:"USESSL" env-default:"false"`
		BucketName      string `env:"BUCKET_NAME" env-default:"general"`
	}
}

func LoadConfig() (*Config, error) {
	config := Config{}
	if err := cleanenv.ReadEnv(&config); err != nil {
		return nil, err
	}

	if config.Env == "dev" {
		if err := cleanenv.ReadConfig(".env", &config); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
