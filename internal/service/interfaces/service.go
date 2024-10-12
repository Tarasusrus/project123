package interfaces

import (
	"BaseApi/internal/models"
	"context"
)

type MusicService interface {
	UserService() UserService
	AdminService() AdminService
}

type UserService interface {
	// Получение текста песни с пагинацией по куплетам
	GetSongText(ctx context.Context, songID int, verse int) (string, error)
	// Получение данных библиотеки с фильтрацией и пагинацией
	GetLibrary(ctx context.Context, filter models.LibraryFilter, page, pageSize int) ([]models.Song, error)
}

type AdminService interface {
	// Удаление песни
	DeleteSong(ctx context.Context, songID int) error
	// Обновление данных песни
	UpdateSong(ctx context.Context, songID int, songData models.SongUpdate) error
	// Добавление новой песни
	AddSong(ctx context.Context, songData models.NewSong) error
}
