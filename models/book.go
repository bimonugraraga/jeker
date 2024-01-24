package models

import "time"

type AllBookStatus struct {
	BookStatusDraft     string
	BookStatusOngoing   string
	BookStatusCompleted string
	BookStatusPublished string
}

func filler() AllBookStatus {
	var status AllBookStatus
	status.BookStatusDraft = "draft"
	status.BookStatusOngoing = "ongoing"
	status.BookStatusCompleted = "completed"
	status.BookStatusPublished = "published"
	return status
}

var BookStatus = filler()

type Book struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	UserId     int       `json:"user_id"`
	CategoryId int       `json:"category_id"`
	Title      *string   `json:"title"`
	Synopsis   *string   `json:"synopsis"`
	Cover      *string   `json:"cover"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	User       User      `gorm:"foreignKey:UserId"`
	Category   Category  `gorm:"foreignKey:CategoryId"`
}
