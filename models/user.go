package models

import (
	"time"
)

type User struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Username  string     `json:"username"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
