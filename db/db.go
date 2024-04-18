package db

import (
	"fmt"
	"github.com/tinarao/url-shortener-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to pgsql:", err)
	}

	log.Println("Connected to pgsql")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations...")
	db.AutoMigrate(&models.Link{})

	DB = DbInstance{
		Db: db,
	}
}
