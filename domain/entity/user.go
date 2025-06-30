package entity

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/model"
)

type User struct {
	ID     uuid.UUID
	Name   string
	Email  string
	Status int64
}

func NewUser(user *model.User) *User {
	return &User{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
