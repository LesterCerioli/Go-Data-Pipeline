package initializers

import (
	"go-pipeline/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.Pipeline{})
}
