package service

import (
	"AuthService/internal/application/services"
	"AuthService/internal/domain/entities"
	"AuthService/internal/infrastructure/persistence/repositories"
	"AuthService/pkg/logger"
	"context"
	"gorm.io/gorm"
)

type authServiceImpl struct {
	accountGen repositories.GenericRepository[entities.Account]
	db         *gorm.DB
}

func NewAuthService(db *gorm.DB, accountGen repositories.GenericRepository[entities.Account]) services.AuthService {
	return &authServiceImpl{
		db:         db,
		accountGen: accountGen,
	}
}

func (a *authServiceImpl) GetAllAccounts() ([]entities.Account, error) {
	ctx := context.Background()
	accounts, _, err := a.accountGen.GetAll(ctx, nil, []string{"Roles"}, 0, 0)
	if err != nil {
		log := logger.GetLogger()
		log.Error("Error fetching accounts: " + err.Error())
		return nil, err
	}
	return accounts, nil
}

func (a *authServiceImpl) Login(username, password string) (*entities.Account, error) {
	// Assuming username is email
	//account, err := a.db.FindByEmail(username)
	//if err != nil {
	//	return nil, err
	//}
	//// TODO: Replace with real password hash check
	//if account.PasswordHash != password {
	//	return nil, errors.New("invalid credentials")
	//}
	return nil, nil
}

func (a *authServiceImpl) CreateAccount(account *entities.Account) (*entities.Account, error) {
	// You may want to hash the password here
	// TODO: Add password hashing
	//created, err := a.accountRepo.Save(account)
	//if err != nil {
	//	return nil, err
	//}
	//return created, nil
	panic("not implemented")
}
