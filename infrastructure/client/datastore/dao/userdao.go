package dao

import (
	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/domain/repository"
	"gorm.io/gorm"
)

// UsersDao struct
type UsersDao struct{}

// NewUsersDao creates a new UserRepository
//
//nolint:ireturn
func NewUsersDao() repository.UserRepository {
	return &UsersDao{}
}

// GetAll is get all
func (u *UsersDao) GetAll(db *gorm.DB) ([]model.User, error) {
	users := make([]model.User, 0)
	err := db.Where("status IN (?)", []model.UserStatus{model.UserStatusActive}).Order("user_id desc").Find(&users).Error
	return users, err
}

// GetAllWithInActive is get all
func (u *UsersDao) GetAllWithInActive(db *gorm.DB) ([]model.User, error) {
	users := make([]model.User, 0)
	err := db.Where("status IN (?)", []model.UserStatus{model.UserStatusActive, model.UserStatusInActive}).Order("user_id desc").Find(&users).Error
	return users, err
}

// Get is get by id
func (u *UsersDao) Get(db *gorm.DB, userID uint64) (model.User, error) {
	user := model.User{}
	err := db.Where("status = ? AND user_id = ?", model.UserStatusActive, userID).First(&user).Error
	return user, err
}

// GetByIDAndEmail is get by id and email
func (u *UsersDao) GetByIDAndEmail(db *gorm.DB, userID uint64, email string) (model.User, error) {
	user := model.User{}
	err := db.Where("status = ? AND user_id = ? AND email = ?", model.UserStatusActive, userID, email).First(&user).Error
	return user, err
}

// GetByEmail is get by  email
func (u *UsersDao) GetByEmail(db *gorm.DB, email string) (model.User, error) {
	user := model.User{}
	err := db.Where("status = ? AND email = ?", model.UserStatusActive, email).First(&user).Error
	return user, err
}

// Create is create new user
func (u *UsersDao) Create(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

// Update is update exist user
func (u *UsersDao) Update(db *gorm.DB, user *model.User) error {
	return db.Model(&model.User{}).Where(
		"status = ? AND user_id = ?",
		model.UserStatusActive,
		user.UserID,
	).Updates(user).Error
}

// InActive is update exist user
func (u *UsersDao) InActive(db *gorm.DB, userID uint64) error {
	return db.Model(&model.User{}).Where(
		"status = ? AND user_id = ?",
		model.UserStatusActive,
		userID,
	).Update("status", model.UserStatusInActive).Error
}
