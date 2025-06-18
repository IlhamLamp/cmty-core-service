package database

import (
	"log"

	"github.com/IlhamLamp/cmty-project-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=projects port=5432 sslmode=disable password=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Project{}, &models.Address{}, &models.Member{}, &models.Role{}, &models.Tag{})
}
