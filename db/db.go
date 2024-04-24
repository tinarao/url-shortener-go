package db

import (
	"fmt"
	"github.com/tinarao/url-shortener-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func Connect() {

	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to pgsql:", err)
		os.Exit(1)
	}

	slog.Info("Connected to pgsql")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Running migrations...")
	if err := db.AutoMigrate(&models.Link{}); err != nil {
		slog.Error("Failed to run auto migrations", err)
		os.Exit(1)
	}

	DB = DbInstance{
		Db: db,
	}
}
