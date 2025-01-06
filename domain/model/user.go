package model

import "time"

type UserStatus int64

const (
	UserStatusActive UserStatus = iota
	UserStatusInActive
)

// User entity
type User struct {
	UserID    uint64    `gorm:"primary_key" json:"user_id" v-put:"omitempty" v-post:"omitempty"`
	Name      string    `json:"name" v-put:"required" v-post:"required"`
	Email     string    `gorm:"index:email" json:"email" v-put:"required" v-post:"required"`
	Password  string    `json:"-" v-put:"required" v-post:"required"`
	Salt      string    `json:"-" v-put:"omitempty" v-post:"omitempty"`
	Status    int64     `gorm:"index:status" json:"status" v-put:"required" v-post:"required"`
	CreatedAt time.Time `json:"_,omitempty"`
	UpdatedAt time.Time `json:"_,omitempty"`
}
