package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"type:varchar(40); unique; not null" binding:"required" json:"name"`
	Email     string    `gorm:"type:varchar(100); unique; not null" binding:"required" json:"email"`
	Password  string    `gorm:"type:varchar(70); not null" json:"password" binding:"required"`
	Point     int       `json:"point" gorm:"type:int"`
}