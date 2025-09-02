package entities

import (
	"github.com/google/uuid"
	"time"
)

type AccountRole struct {
	AccountID    uuid.UUID  `gorm:"type:uuid;primaryKey"`
	RoleID       string     `gorm:"primaryKey;size:20;column:role_id"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID"`
	AssignedTime time.Time  `gorm:"autoCreateTime;column:assigned_time"`
	AssignedBy   *uuid.UUID `gorm:"type:uuid;column:assigned_by"`
	EffectDate   *time.Time `gorm:"column:effect_date"`
	EndDate      *time.Time `gorm:"column:end_date"`
}
