package dao

import (
	"github.com/google/uuid"
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/domain/repository"
	"gorm.io/gorm"
)

// UserDao struct.
type UserDao struct{}

// NewUserDao creates a new UserRepository.
func NewUserDao() repository.UserRepository {
	return &UserDao{}
}

// GetAll is get all.
func (u *UserDao) GetAll(db *gorm.DB) ([]model.User, error) {
	users := make([]model.User, 0)
	err := db.Where("status IN (?)", []model.UserStatus{model.UserStatusActive}).Where("deleted_at IS NULL").Order("id desc").Find(&users).Error
	return users, err
}

// GetAllWithInActive is get all.
func (u *UserDao) GetAllWithInActive(db *gorm.DB) ([]model.User, error) {
	users := make([]model.User, 0)
	err := db.Where("status IN (?)", []model.UserStatus{model.UserStatusActive, model.UserStatusInActive}).Where("deleted_at IS NULL").Order("id desc").Find(&users).Error
	return users, err
}

// Get is get by id.
func (u *UserDao) Get(db *gorm.DB, userID uuid.UUID) (model.User, error) {
	user := model.User{}
	err := db.Where("status = ? AND id = ?", model.UserStatusActive, userID).Where("deleted_at IS NULL").First(&user).Error
	return user, err
}

// GetByIDAndEmail is get by id and email.
func (u *UserDao) GetByIDAndEmail(db *gorm.DB, userID uuid.UUID, email string) (model.User, error) {
	user := model.User{}
	err := db.Where("status = ? AND id = ? AND email = ?", model.UserStatusActive, userID, email).Where("deleted_at IS NULL").First(&user).Error
	return user, err
}

// GetByEmail is get by  email.
func (u *UserDao) GetByEmail(db *gorm.DB, email string) (model.User, error) {
	user := model.User{}
	err := db.Where("status = ? AND email = ?", model.UserStatusActive, email).Where("deleted_at IS NULL").First(&user).Error
	return user, err
}

// Create is create new user.
func (u *UserDao) Create(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

// Update is update exist user.
func (u *UserDao) Update(db *gorm.DB, user *model.User) error {
	return db.Model(&model.User{}).Where(
		"status = ? AND id = ?",
		model.UserStatusActive,
		user.ID,
	).Updates(user).Error
}

// InActive is update exist user.
func (u *UserDao) InActive(db *gorm.DB, userID uuid.UUID) error {
	return db.Model(&model.User{}).Where(
		"status = ? AND id = ?",
		model.UserStatusActive,
		userID,
	).Update("status", model.UserStatusInActive).Error
}
