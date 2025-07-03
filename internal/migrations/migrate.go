package migrations

import (
	"gorm.io/gorm"
	"microservice/internal/models"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Token{},
	)
}
