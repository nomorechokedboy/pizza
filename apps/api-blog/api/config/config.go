package config

import (
	"github.com/spf13/viper"
)

type Config struct {
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
}

func LoadConfig() (*Config, error) {
	config := new(Config)
	env := viper.SetDefault

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	env("database.db_port", "5432")
	env("database.db_host", "localhost")
	env("database.db_name", "pizza")
	env("database.db_user", "postgres")
	env("database.db_password", "postgres")
	env("server.host", "")
	env("server.port", "8080")
	env("authconfig.jwt_refresh_secret", "refresh-secret")
	env("authconfig.jwt_secret", "my-secret")
	env("authemail.email", "kristiannguyen276@gmail.com")
	env("authemail.email_password", "figjbdfsggwhcvbr")
	env("appapi.api_link", "Hello")
	env("minio.end_point", "localhost:9000")
	env("minio.accesskeyid", "admin")
	env("minio.secret_access_key", "admin123")
	env("minio.usessl", "false")
	env("minio.bucket_name", "general")
	env("audioapi.link", viper.Get("audio_link"))
	env("audioapi.key", viper.Get("audio_key"))

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
