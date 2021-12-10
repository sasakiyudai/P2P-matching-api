package model

import (
	"time"
)

type Product struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index"`
	Name      string     `gorm:"type:varchar(40); not null" binding:"required" json:"name"`
	Comment   string     `gorm:"type:varchar(150); not null" json:"comment"`
	Price     int        `json:"price" gorm:"type:int; not null"`
	UserID    uint       `gorm:"not null"`
	User      User
}
