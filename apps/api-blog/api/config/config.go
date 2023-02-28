package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
