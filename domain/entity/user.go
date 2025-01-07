package entity

import "github.com/howood/moggiecollector/domain/model"

type User struct {
	UserID uint64
	Name   string
	Email  string
	Status int64
}

func NewUser(user *model.User) *User {
	return &User{
		UserID: user.UserID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
}
