package service

import (
	"gin/internal/domain/model"
	"gin/internal/domain/repository"
	"gin/internal/interface/dto/http"
)

type UserService struct {
	repository repository.UserRepository
}

func (u UserService) CreateUser(request http.UserCreateRequest) (http.UserDetailsResponse, error) {
	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		return http.UserDetailsResponse{}, err
	}
	user := model.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: hashedPassword,
	}
	err = u.repository.Save(&user)
	if err != nil {
		return http.UserDetailsResponse{}, err
	}
	return http.UserDetailsResponse{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
