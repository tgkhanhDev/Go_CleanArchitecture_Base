package services

import (
	"AuthService/internal/application/dtos/request"
	"AuthService/internal/domain/entities"
)

type UserService interface {
	GetUserByID(id int8) (*entities.User, error)
	GetAllUsers() ([]*entities.User, error)
	CreateUser(user request.UserCreateRequest) (*entities.User, error)
}
