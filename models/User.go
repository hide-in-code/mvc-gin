package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id int64 `json:"id"` 
	CreatedAt time.Time `json:"created_at"` 
	UpdatedAt time.Time `json:"updated_at"` 
	DeletedAt time.Time `json:"deleted_at"` 
	Username string `json:"username"` 
	Password string `json:"password"` 
	Role int64 `json:"role"` 
}
