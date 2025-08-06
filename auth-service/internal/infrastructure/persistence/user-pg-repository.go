package persistence

import (
	"database/sql"
	"gin/internal/domain/model"
	"gin/internal/domain/repository"
)

type PgUserRepository struct {
	// db *gorm.DB // Uncomment if you need a database connection
	DB *sql.DB
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
