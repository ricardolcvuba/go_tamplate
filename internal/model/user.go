package model

import (
	"time"

	"github.com/google/uuid"
)

// User estructura que representa a un usuario en la base de datos
type User struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey"`
	Name      string     `gorm:"not null;size:25" json:"name"`
	Email     string     `gorm:"uniqueIndex;not null" json:"email"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
