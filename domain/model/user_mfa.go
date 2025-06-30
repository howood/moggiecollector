package model

import (
	"github.com/google/uuid"
)

type MfaType string

const (
	MfaTypeTOTP     MfaType = "TOTP"
	MfaTypeWebAuthn MfaType = "WebAuthn"
)

// UserMfa entity.
type UserMfa struct {
	BaseModel

	UserID    uuid.UUID `gorm:"type:uuid;index:user_mfas_user_id_index;size:255;not null"`
	MfaType   MfaType   `gorm:"type:varchar(50);index:user_mfas_mfa_type_index;not null"`
	Secret    string    `gorm:"type:varchar(255)"`
	IsDefault bool      `gorm:"type:boolean;default:false"`
}
