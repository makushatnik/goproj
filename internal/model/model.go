package model

import (
	"fmt"

	"gorm.io/gorm"
)

func InitModel(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Card{})
	if err != nil {
		return fmt.Errorf("failed to init models: %w", err)
	}

	return nil
}
