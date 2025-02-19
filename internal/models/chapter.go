package models

import "time"

type Chapter struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Number       int    `gorm:"not null" json:"number"`
	Title        string `gorm:"not null" json:"title"`
	EnglishTitle string `gorm:"column:english_title" json:"english_title"`
	// PaliTitle string `gorm:"column:pali_title" json:"pali_title"`
	VersesCount int       `gorm:"column:verses_count" json:"verses_count"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
	Verses      []Verse   `gorm:"foreignKey:ChapterID" json:"verses,omitempty"`
}
