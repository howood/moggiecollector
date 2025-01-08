package response

import "github.com/google/uuid"

type UserResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Status int64     `json:"status"`
}
