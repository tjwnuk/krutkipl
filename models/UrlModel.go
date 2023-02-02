package models

import "time"

type UrlModel struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	OriginalURL  string
	Token        string
	ShortenedURL string
}
