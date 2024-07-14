package entity

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json:"role" gorm:"type:enum('admin','user')"`
	ImageProfile string `json:"image_profile"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
