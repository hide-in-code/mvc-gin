package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int64     `json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;default:null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:null" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:null" json:"deleted_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      int64     `json:"role"`
}
