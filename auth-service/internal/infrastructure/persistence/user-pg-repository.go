package persistence

import (
	repository "gin/internal/application/interface/repositories"
	model "gin/internal/domain/entities"
	"gorm.io/gorm"
)

type PgUserRepository struct {
	dbOrm *gorm.DB
	//db    *sql.DB
}

func (p PgUserRepository) Save(user *model.User) (*model.User, error) {
	if err := p.dbOrm.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (p PgUserRepository) GetById(id int64) (*model.User, error) {
	var user model.User
	result := p.dbOrm.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // User not found
		}
		return nil, result.Error // Other error
	}
	panic("implement me")
}

var _ repository.UserRepository = &PgUserRepository{}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &PgUserRepository{
		dbOrm: db,
	}
}
