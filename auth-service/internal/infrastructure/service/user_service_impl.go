package service

import (
	"AuthService/internal/application/dtos/request"
	"AuthService/internal/application/services"
	"AuthService/internal/domain/entities"
	"AuthService/internal/infrastructure/persistence/repositories"
)

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) services.UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (u *userServiceImpl) GetUserByID(id int8) (*entities.User, error) {
	user, err := u.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userServiceImpl) GetAllUsers() ([]*entities.User, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userServiceImpl) CreateUser(req request.UserCreateRequest) (*entities.User, error) {
	return nil, nil
}
