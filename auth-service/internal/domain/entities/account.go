package entities

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid();column:id"`
	Email        string    `gorm:"size:255;unique;not null;column:email"`
	PasswordHash string    `gorm:"size:255;not null;column:password_hash"`
	IsActive     bool      `gorm:"default:true;column:is_active"`
	Role         string    `gorm:"size:20;column:role"`
	CreatedAt    time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at"`
}
