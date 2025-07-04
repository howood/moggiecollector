package response

import "github.com/google/uuid"

type UserResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Status int64     `json:"status"`
}
