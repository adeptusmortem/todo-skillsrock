package database

import (
	"todo-list/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(sConnection string) error {
	db, err := gorm.Open(postgres.Open(sConnection), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&models.Task{},
	)
}