package entities

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"size:100;not null;unique"`
	PasswordHash string    `gorm:"size:255;not null"`
	Email        string    `gorm:"size:150;unique"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
