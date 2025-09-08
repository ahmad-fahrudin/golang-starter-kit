package database

import (
	"log"

	"golang-starter-kit/config"
	migr "golang-starter-kit/database/migrations"
	"golang-starter-kit/database/seeders"
	"golang-starter-kit/pkg/utils"

	"gorm.io/gorm"
)

func Seed(cfg *config.Config, db *gorm.DB) error {
	// Run migrations
	if err := migr.MigrateUp(db); err != nil {
		return err
	}

	// Run seeders
	if err := seeders.SeedUsers(db); err != nil {
		return err
	}

	log.Println("Database migrated and seeded successfully")
	_ = utils.HashPassword // ensure utils imported when not used elsewhere
	return nil
}

// MigrateOnly runs migrations without running seeders
func MigrateOnly(db *gorm.DB) error {
	if err := migr.MigrateUp(db); err != nil {
		return err
	}
	log.Println("Database migrated successfully (migrate only)")
	return nil
}

// SeedOnly runs seeders without running migrations
func SeedOnly(db *gorm.DB) error {
	if err := seeders.SeedUsers(db); err != nil {
		return err
	}
	log.Println("Database seeded successfully (seed only)")
	return nil
}
