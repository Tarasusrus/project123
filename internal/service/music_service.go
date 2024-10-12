package service

import (
	"BaseApi/internal/service/interfaces"
	"BaseApi/internal/service/repository"
)

type musicService struct {
	userService  interfaces.UserService
	adminService interfaces.AdminService
}

func (m *musicService) UserService() interfaces.UserService {
	return m.userService
}

func (m *musicService) AdminService() interfaces.AdminService {
	return m.adminService
}

func NewMusicService(repository repository.Repository) interfaces.MusicService {
	return &musicService{
		userService:  NewUserService(repository),
		adminService: NewAdminService(repository),
	}
}
