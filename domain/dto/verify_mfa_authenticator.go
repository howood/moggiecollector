package dto

import "github.com/google/uuid"

type VerifyMfaAuthenticator struct {
	UserID   uuid.UUID
	Passcode string
}
