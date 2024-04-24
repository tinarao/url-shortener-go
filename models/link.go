package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	OriginalLink string `json:"original_link"`
	Alias        string `json:"alias" gorm:"uniqueIndex"`
}
