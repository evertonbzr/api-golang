package api

import (
	"fmt"
	"log/slog"
)

type APIConfig struct {
	Port string
}

func Run(cfg *APIConfig) {
	slog.Info(fmt.Sprintf("API running on port %s", cfg.Port))
}
