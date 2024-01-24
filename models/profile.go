package models

import "time"

type Profile struct {
	ID              int        `json:"id" gorm:"primaryKey"`
	UserId          int        `json:"user_id"`
	ProfilePictures *string    `json:"profile_pictures"`
	Bio             *string    `json:"bio"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	User            User       `gorm:"foreignKey:UserId"`
}
