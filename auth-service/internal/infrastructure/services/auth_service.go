package service

import (
	repository "gin/internal/application/interface/repositories"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a AuthService) Login(username, password string) (string, error) {
	//TODO implement me
	return "Hello " + username, nil

	panic("implement me")
}
