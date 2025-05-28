package db

import (
	"github.com/Deepankar20/shinra/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=probo_clone port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the User schema
	err = db.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Event{},
		&models.Trade{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
