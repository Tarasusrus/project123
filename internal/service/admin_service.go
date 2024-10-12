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
	return a.Repository.DeleteSong(ctx, songID)
}

func (a adminService) UpdateSong(ctx context.Context, songID int, songData models.SongUpdate) error {
	return a.Repository.UpdateSong(ctx, songID, songData)
}

func (a adminService) AddSong(ctx context.Context, songData models.NewSong) error {
	return a.Repository.AddSong(ctx, songData)
}
