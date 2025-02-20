package models

import "time"

type Verse struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	ChapterID         uint      `gorm:"index" json:"chapter_id"`
	VerseNumber       int       `gorm:"column:verse_number" json:"verse_number"`
	Text              string    `json:"text"`
	EnglishText       string    `json:"english_text"`
	StoryTitle        string    `gorm:"column:story_title" json:"story_title"`
	EnglishStoryTitle string    `gorm:"column:english_story_title" json:"english_story_title"`
	Story             string    `json:"story"`
	EnglishStory      string    `gorm:"column:english_story" json:"english_story"`
	CreatedAt         time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"not null" json:"updated_at"`
	Chapter           Chapter   `gorm:"foreignKey:ChapterID" json:"chapter,omitempty"`
}
