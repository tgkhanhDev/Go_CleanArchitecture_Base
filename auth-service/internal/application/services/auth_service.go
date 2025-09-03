package services

import (
	"AuthService/internal/application/dtos/request"
	"AuthService/internal/application/dtos/response"
	"AuthService/internal/domain/entities"
)

type AuthService interface {
	GetAllAccounts() ([]entities.Account, error) //test purpose only
	Login(request request.LoginRequest) (*response.LoginResponse, error)
	CreateAccount(account request.CreateAccountRequest) (*entities.Account, error)
}
