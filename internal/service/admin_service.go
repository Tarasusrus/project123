package service

import (
	"BaseApi/internal/models"
	"BaseApi/internal/service/interfaces"
	"BaseApi/internal/service/repository"
	"context"
)

func NewAdminService(repository repository.Repository) interfaces.AdminService {
	return &adminService{Repository: repository}
}

type adminService struct {
	Repository repository.Repository
}

func (a adminService) DeleteSong(ctx context.Context, songID int) error {
	//TODO implement me
	panic("implement me")
}

func (a adminService) UpdateSong(ctx context.Context, songID int, songData models.SongUpdate) error {
	//TODO implement me
	panic("implement me")
}

func (a adminService) AddSong(ctx context.Context, songData models.NewSong) error {
	//TODO implement me
	panic("implement me")
}
