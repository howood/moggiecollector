package dao

import (
	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserMfaDao struct.
type UserMfaDao struct{}

// NewUserMfaDao creates a new UserMfaRepository.
func NewUserMfaDao() repository.UserMfaRepository {
	return &UserMfaDao{}
}

func (u *UserMfaDao) Get(db *gorm.DB, userID uuid.UUID, mfaType model.MfaType) (*model.UserMfa, error) {
	userMfa := model.UserMfa{}
	err := db.Where("user_id = ? AND mfa_type = ?", userID, mfaType).First(&userMfa).Error
	if err != nil {
		return nil, err
	}
	return &userMfa, nil
}

func (u *UserMfaDao) GetDefault(db *gorm.DB, userID uuid.UUID) (*model.UserMfa, error) {
	userMfa := model.UserMfa{}
	err := db.Where("user_id = ? AND is_default = ?", userID, true).First(&userMfa).Error
	if err != nil {
		return nil, err
	}
	return &userMfa, nil
}

// UnsetDefault is unset default mfa type
func (u *UserMfaDao) UnsetDefault(db *gorm.DB, userID uuid.UUID, excludeMfaType model.MfaType) error {
	return db.Model(&model.UserMfa{}).Where("user_id = ? AND mfa_type <> ?", userID, excludeMfaType).Update("is_default", false).Error
}

// Update is update exist user
func (u *UserMfaDao) Upsert(db *gorm.DB, userMfa *model.UserMfa) error {
	return db.Omit(clause.Associations).Save(userMfa).Error
}
