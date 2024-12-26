package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port string `env:"PORT,required"`
}

func Parse() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("failed to load .env file", logrus.WithError(err))
	}

	var cfg Config

	err = env.Parse(&cfg)
	if err != nil {
		logrus.Warn("failed to parse envs", logrus.WithError(err))
		return nil, err
	}

	return &cfg, err
}
