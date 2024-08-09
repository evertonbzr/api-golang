package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/evertonbzr/api-golang/internal/config"
	"github.com/evertonbzr/api-golang/internal/database"
	"github.com/evertonbzr/api-golang/internal/model"
)

func main() {
	config.Load(os.Getenv("ENV"))

	db := database.InitDatabase(config.DATABASE_URL)
	slog.Info("Connected to database")

	err := db.AutoMigrate(&model.User{}, &model.Book{}, &model.Borrowing{})

	if err != nil {
		log.Fatal("Error migrating database", "error", err)
	}

	slog.Info("Database migrated")
}
