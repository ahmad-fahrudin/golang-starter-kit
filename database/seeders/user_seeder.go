package seeders

import (
	"log"
	"time"

	"golang-starter-kit/internal/models"
	"golang-starter-kit/pkg/utils"

	"gorm.io/gorm"
)

// SeedUsers inserts initial users into the database
func SeedUsers(db *gorm.DB) error {
	// check if users already exist
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		log.Println("users already seeded, skipping")
		return nil
	}

	hashed, err := utils.HashPassword("password123")
	if err != nil {
		return err
	}

	users := []models.User{
		{
			Name:      "Admin",
			Email:     "admin@example.com",
			Password:  hashed,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return db.Create(&users).Error
}
