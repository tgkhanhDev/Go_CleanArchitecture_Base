package repositories

import (
	model "gin/internal/domain/entities"
)

type UserRepository interface {
	GetById(id int64) (*model.User, error)
	Save(user *model.User) (*model.User, error)
}
