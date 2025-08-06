package repository

import (
	"gin/internal/domain/model"
)

type UserRepository interface {
	GetById(id int64) (*model.User, error)
	Save(user *model.User) error
}
