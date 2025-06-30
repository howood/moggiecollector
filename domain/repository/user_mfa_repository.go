package repository

import (
	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/model"
	"gorm.io/gorm"
)

// UserRepository interface
type UserMfaRepository interface {
	GetAll(db *gorm.DB) ([]model.User, error)
	GetAllWithInActive(db *gorm.DB) ([]model.User, error)
	Get(db *gorm.DB, userID uuid.UUID) (model.User, error)
	GetByIDAndEmail(db *gorm.DB, userID uuid.UUID, email string) (model.User, error)
	GetByEmail(db *gorm.DB, email string) (model.User, error)
	Create(db *gorm.DB, user *model.User) error
	Update(db *gorm.DB, user *model.User) error
	InActive(db *gorm.DB, userID uuid.UUID) error
}
