package models

import (
	"flag"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Configs struct {
	ID        int64     `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

var dns = flag.String("dns", "host=localhost user=dracuxan password=pass dbname=blodbank port=5432", "Postgres connection URL")

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(*dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Configs{})
	return db, err
}
