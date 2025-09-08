package migrations

import (
	"time"

	"gorm.io/gorm"
)

// Users migration - GORM will use this struct shape only for migration
type Users struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// MigrateUp runs the migration for users table
func MigrateUp(db *gorm.DB) error {
	return db.AutoMigrate(&Users{})
}

// MigrateDown drops the users table
func MigrateDown(db *gorm.DB) error {
	return db.Migrator().DropTable(&Users{})
}
