package service

import (
	request "gin/internal/application/dtos/request"
	response "gin/internal/application/dtos/response"
	repository "gin/internal/application/interface/repositories"
	model "gin/internal/domain/entities"
	pwd "gin/pkg/password"
)

type UserService struct {
	repository repository.UserRepository
}

func (u UserService) CreateUser(request request.UserCreateRequest) (response.UserDetailsResponse, error) {
	hashedPassword, err := pwd.HashPassword(request.Password)
	if err != nil {
		return response.UserDetailsResponse{}, err
	}
	user := model.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: hashedPassword,
	}
	userCreated, err := u.repository.Save(&user)
	if err != nil {
		return response.UserDetailsResponse{}, err
	}
	return response.UserDetailsResponse{
		Username: userCreated.Username,
		Email:    userCreated.Email,
	}, nil
}
