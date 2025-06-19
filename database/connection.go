package database

import (
	"fmt"
	"log"

	"github.com/IlhamLamp/cmty-project-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConnection struct {
	Host     string
	Name     string
	User     string
	Password string
}

func Connect(cfg DbConnection) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s port=5432 sslmode=disable password=%s",
		cfg.Host, cfg.Name, cfg.User, cfg.Password,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Project{}, &models.Address{}, &models.Member{}, &models.Role{}, &models.Tag{})
}
