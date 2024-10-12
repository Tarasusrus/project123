package service

import (
	"BaseApi/internal/models"
	"BaseApi/internal/service/interfaces"
	"BaseApi/internal/service/repository"
	"context"
)

type userService struct {
	Repository repository.Repository
}

func NewUserService(repository repository.Repository) interfaces.UserService {
	return &userService{Repository: repository}
}

func (u userService) GetSongText(ctx context.Context, songID int, verse int) (string, error) {
	return u.Repository.GetSongText(ctx, songID, verse)
}

func (u userService) GetLibrary(ctx context.Context, filter models.LibraryFilter, page, pageSize int) ([]models.Song, error) {
	return u.Repository.GetLibrary(ctx, filter, page, pageSize)
}
