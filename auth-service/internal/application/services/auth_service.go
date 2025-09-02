package services

import "AuthService/internal/domain/entities"

type AuthService interface {
	GetAllAccounts() ([]entities.Account, error) //test purpose only
	Login(username, password string) (*entities.Account, error)
	CreateAccount(account *entities.Account) (*entities.Account, error)
}
