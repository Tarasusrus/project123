package repository

import (
	"BaseApi/internal/models"
	"context"
)

type Repository interface {
	GetSongText(ctx context.Context, songID int, verse int) (string, error)
	GetLibrary(ctx context.Context, filter models.LibraryFilter, page, pageSize int) ([]models.Song, error)
	DeleteSong(ctx context.Context, songID int) error
	UpdateSong(ctx context.Context, songID int, songData models.SongUpdate) error
	AddSong(ctx context.Context, songData models.NewSong) error
}
