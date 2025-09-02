package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid();column:id"`
	AccountID uuid.UUID `gorm:"type:uuid;not null;unique;column:account_id"`
	FullName  string    `gorm:"size:255;not null;column:full_name"`
	Phone     string    `gorm:"size:50"`
	AvatarUrl string    `gorm:"size:255;column:avatar_url"`
	Bio       string    `gorm:"size:255;column:bio"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at"`
}
