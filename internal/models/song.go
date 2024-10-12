package models

import _ "gorm.io/gorm"

type LibraryFilter struct {
	Group string
	Song  string
}

type NewSong struct {
	Group string `json:"group" gorm:"not null"`
	Song  string `json:"song" gorm:"not null"`
}

type Song struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Group       string `gorm:"not null" json:"group"`
	Song        string `gorm:"not null" json:"song"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongUpdate struct {
	Group       *string `json:"group,omitempty"`
	Song        *string `json:"song,omitempty"`
	ReleaseDate *string `json:"release_date,omitempty"`
	Text        *string `json:"text,omitempty"`
	Link        *string `json:"link,omitempty"`
}
