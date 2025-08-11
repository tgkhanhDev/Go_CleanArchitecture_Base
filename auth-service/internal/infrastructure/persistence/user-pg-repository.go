package persistence

import (
	"database/sql"
	repository "gin/internal/application/interface/repositories"
	model "gin/internal/domain/entities"
)

type PgUserRepository struct {
	// db *gorm.DB // Uncomment if you need a databases connection
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &PgUserRepository{
		DB: db,
	}
}

func (p PgUserRepository) GetById(id int64) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p PgUserRepository) Save(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

var _ repository.UserRepository = &PgUserRepository{}
