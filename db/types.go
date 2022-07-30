package db

import "gorm.io/gorm"

// Handler is handler
type Handler struct {
	DB *gorm.DB
}
