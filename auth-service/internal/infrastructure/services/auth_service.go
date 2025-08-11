package service

import (
	request "gin/internal/application/dtos/request"
	http "gin/internal/application/dtos/response"
	repository "gin/internal/application/interface/repositories"
	model "gin/internal/domain/entities"
	pwd "gin/pkg/password"
)

type AuthService struct {
	repo repository.UserRepository
}

func (a AuthService) RegisterUser(registerReq request.RegisterRequest) (http.UserDetailsResponse, error) {
	hashedPwd, err := pwd.HashPassword(registerReq.Password)

	user := &model.User{
		Username:     registerReq.Username,
		PasswordHash: hashedPwd,
		Email:        registerReq.Email,
	}
	user, err = a.repo.Save(user)

	return http.Of(user), err
	//TODO implement me
	panic("implement me")
}

func (a AuthService) Login(username, password string) (string, error) {
	//TODO implement me
	return "Hello " + username, nil

	panic("implement me")
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
