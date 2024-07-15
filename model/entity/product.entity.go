package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
