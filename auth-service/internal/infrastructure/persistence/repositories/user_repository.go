package repositories

import (
	"AuthService/internal/domain/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (uRepo UserRepository) FindAll() ([]*entities.User, error) {
	var users []*entities.User
	if err := uRepo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (uRepo UserRepository) FindById(id int8) (*entities.User, error) {
	var user entities.User
	rs := uRepo.db.First(&user, "id = ?", id)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return &user, nil
}

func (uRepo UserRepository) FindByEmail(email string) (*[]entities.User, error) {
	var user []entities.User
	rs := uRepo.db.Find(&user, "email = ?", email)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return &user, nil
}

func (uRepo UserRepository) Save(user *entities.User) (*entities.User, error) {
	tx := uRepo.db.Save(user)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return nil, tx.Error
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
