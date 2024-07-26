package config

import (
	"log"
	"log/slog"

	"github.com/spf13/viper"
)

var (
	PORT      string
	ENV       string
	NAME      string
	REDIS_URL string
)

func Load(env string) {
	slog.Info("Loading config...", "env", env)

	if env == "" {
		viper.Set("ENV", "development")
		viper.Set("PORT", "8081")
		viper.Set("NAME", "api-golang")
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error loading .env file", "error", err)
		}
	} else {
		viper.AutomaticEnv()
	}

	PORT = viper.GetString("PORT")
	ENV = viper.GetString("ENV")
	NAME = viper.GetString("NAME")
	REDIS_URL = viper.GetString("REDIS_URL")
}

func IsDevelopment() bool {
	return ENV == "development"
}

func IsProduction() bool {
	return ENV == "production"
}

func IsTest() bool {
	return ENV == "test"
}
