package model

import "time"

type UserStatus int64

const (
	UserStatusActive UserStatus = iota
	UserStatusInActive
)

// User entity
type User struct {
	UserID    uint64 `gorm:"primary_key"  json:"user_id"     v-post:"omitempty" v-put:"omitempty"`
	Name      string `json:"name"         v-post:"required"  v-put:"required"`
	Email     string `gorm:"index:email"  json:"email"       v-post:"required"  v-put:"required"`
	Password  string `json:"-"            v-post:"required"  v-put:"required"`
	Salt      string `json:"-"            v-post:"omitempty" v-put:"omitempty"`
	Status    int64  `gorm:"index:status" json:"status"      v-post:"required"  v-put:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
