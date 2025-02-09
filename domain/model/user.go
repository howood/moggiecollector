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
	Email     string `gorm:"index:users_email_index"`
	Password  string
	Salt      string
	Status    int64 `gorm:"index:users_status_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index:users_deleted_at_index"`
}
