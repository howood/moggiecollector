package dto

import "github.com/google/uuid"

type UserMfaTotpDto struct {
	UserID    uuid.UUID
	Secret    string
	IsDefault bool
	Passcode  string
}
