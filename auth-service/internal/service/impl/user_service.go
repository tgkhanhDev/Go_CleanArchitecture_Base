package service

import (
	req "gin/internal/dto/request"
	res "gin/internal/dto/response"
	model "gin/internal/models"
	"gin/internal/repository"
	service "gin/internal/util"
)

type UserService struct {
	repository repository.UserRepository
}

func (u UserService) CreateUser(request req.UserCreateRequest) (res.UserDetailsResponse, error) {
	hashedPassword, err := service.HashPassword(request.Password)
	if err != nil {
		return res.UserDetailsResponse{}, err
	}
	user := model.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: hashedPassword,
	}
	err = u.repository.Save(&user)
	if err != nil {
		return res.UserDetailsResponse{}, err
	}
	return res.UserDetailsResponse{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
