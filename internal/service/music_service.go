package service

import "BaseApi/internal/service/interfaces"

type MusicService struct {
	UserService  interfaces.UserService
	AdminService interfaces.AdminService
}
