package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	ID            uint   `json:"id"`
	OriginalLink  string `json:"original_link"`
	ShortenedLink string `json:"shortened_link"`
}
