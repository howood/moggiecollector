package repository

import (
	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/model"
	"gorm.io/gorm"
)

// UserRepository interface.
type UserMfaRepository interface {
	Get(db *gorm.DB, userID uuid.UUID, mfaType model.MfaType) (*model.UserMfa, error)
	GetDefault(db *gorm.DB, userID uuid.UUID) (*model.UserMfa, error)
	UnsetDefault(db *gorm.DB, userID uuid.UUID, excludeMfaType model.MfaType) error
	Upsert(db *gorm.DB, user *model.UserMfa) error
}
