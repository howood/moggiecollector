package model

import (
	"time"

	"github.com/google/uuid"
)

type UserStatus int64

const (
	UserStatusActive UserStatus = iota
	UserStatusInActive
)

// User entity
type User struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey;size:255;default:uuid_generate_v4()"`
	Name      string
	Email     string `gorm:"index:email"`
	Password  string
	Salt      string
	Status    int64 `gorm:"index:status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
