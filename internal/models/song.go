package models

type LibraryFilter struct {
	Group string
	Song  string
}

type NewSong struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type Song struct {
	ID          int    `json:"id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
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
