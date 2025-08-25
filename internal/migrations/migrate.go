package migrations

import (
	"gofax-billing/internal/models"

	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Token{},
		&models.Transaction{},
	)
}
