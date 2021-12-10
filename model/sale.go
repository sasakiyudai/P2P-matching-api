package model

import (
	"time"
)

type Sale struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ProductID uint      `gorm:"not null"`
	Product   Product
	UserID    uint `gorm:"not null"`
	User      User
}
