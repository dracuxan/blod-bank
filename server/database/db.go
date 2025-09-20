package database

import (
	"log"
	"os"

	"github.com/dracuxan/blod-bank/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	db.AutoMigrate(&models.Configs{})

	return db
}
