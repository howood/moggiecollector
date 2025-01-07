package model

import "time"

type UserStatus int64

const (
	UserStatusActive UserStatus = iota
	UserStatusInActive
)

// User entity
type User struct {
	UserID    uint64 `gorm:"primary_key"`
	Name      string
	Email     string `gorm:"index:email"`
	Password  string
	Salt      string
	Status    int64 `gorm:"index:status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
