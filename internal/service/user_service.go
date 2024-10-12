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
	//TODO implement me
	panic("implement me")
}

func (u userService) GetLibrary(ctx context.Context, filter models.LibraryFilter, page, pageSize int) ([]models.Song, error) {
	//TODO implement me
	panic("implement me")
}
