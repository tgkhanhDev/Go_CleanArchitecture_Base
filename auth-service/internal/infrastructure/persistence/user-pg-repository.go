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

func (p PgUserRepository) Save(user *model.User) (*model.User, error) {
	_, err := p.DB.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", user.ID, user.Username, user.Email)
	if err != nil {
		return nil, err
	}

	var savedUser model.User
	err = p.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", user.ID).
		Scan(&savedUser.ID, &savedUser.Username, &savedUser.Email)
	if err != nil {
		return nil, err
	}

	return &savedUser, nil
}

func (p PgUserRepository) GetById(id int64) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

var _ repository.UserRepository = &PgUserRepository{}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &PgUserRepository{
		DB: db,
	}
}
