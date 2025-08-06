package service

import "gin/internal/domain/repository"

type UserService struct {
	repository repository.UserRepository
}
