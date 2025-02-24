package models

type SearchResponse struct {
	Chapters []Chapter `json:"chapters"`
	Verses   []Verse   `json:"verses"`
}
