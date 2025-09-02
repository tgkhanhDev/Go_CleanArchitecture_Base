package repositories

import (
	"AuthService/internal/domain/entities"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Custom methods for Account
func (a *AccountRepository) FindAll() ([]*entities.Account, error) {
	var accounts []*entities.Account
	if err := a.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (a *AccountRepository) FindByEmail(email string) (*entities.Account, error) {
	var account entities.Account
	if err := a.db.Where("email = ?", email).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountRepository) FindByUsername(username string) (*entities.Account, error) {
	// Username field does not exist in Account entity, so this should be implemented if you have a username field or join with User entity.
	return nil, gorm.ErrRecordNotFound
}
