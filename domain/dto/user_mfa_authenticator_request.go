package dto

import "github.com/google/uuid"

type UpsertUserAuthenticator struct {
	UserID    uuid.UUID
	Secret    string
	IsDefault bool
	Passcode  string
}
