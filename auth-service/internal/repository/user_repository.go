package repository

import model "gin/internal/models"

type UserRepository interface {
	GetById(id int64) (*model.User, error)
	Save(user *model.User) error
}
