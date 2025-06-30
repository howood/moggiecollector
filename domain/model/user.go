package model

type UserStatus int64

const (
	UserStatusActive UserStatus = iota
	UserStatusInActive
)

// User entity.
type User struct {
	BaseModel

	Name     string
	Email    string `gorm:"index:users_email_index"`
	Password string
	Salt     string
	Status   int64 `gorm:"index:users_status_index"`
}
