package main

import (
	"log/slog"
	"os"

	"github.com/evertonbzr/api-golang/internal/api"
	"github.com/evertonbzr/api-golang/internal/config"
)

func main() {
	config.Load(os.Getenv("ENV"))

	slog.Info("Starting API...", "port", config.PORT, "env", config.ENV)

	cfg := &api.APIConfig{
		Port: config.PORT,
	}

	api.Run(cfg)
}
