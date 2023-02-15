package utils

import (
	"api/src/config"
	"fmt"
)

func GetDbURI(cfg *config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)
}

func GetServerAddress(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
}

func GetDataTypeAddress[T any](data T) *T {
	return &data
}
