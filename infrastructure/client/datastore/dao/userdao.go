package dao

import (
	"context"

	"github.com/howood/moggiecollector/application/actor"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/infrastructure/client/datastore"
)

// UsersDao struct
type UsersDao struct {
	instance datastore.DatastoreInstance
	ctx      context.Context
}

// NewUsersDao creates a new UserRepository
func NewUsersDao(ctx context.Context, instance datastore.DatastoreInstance) *UsersDao {
	return &UsersDao{ctx: ctx, instance: instance}
}

// GetAll is get all
func (u *UsersDao) GetAll() ([]entity.User, error) {
	users := make([]entity.User, 0)
	err := u.instance.GetClient().Where("status IN (?)", []entity.UserStatus{entity.UserStatusActive}).Order("user_id desc").Find(&users).Error
	return users, err
}

// GetAllWithInActive is get all
func (u *UsersDao) GetAllWithInActive() ([]entity.User, error) {
	users := make([]entity.User, 0)
	err := u.instance.GetClient().Where("status IN (?)", []entity.UserStatus{entity.UserStatusActive, entity.UserStatusInActive}).Order("user_id desc").Find(&users).Error
	return users, err
}

// Get is get by id
func (u *UsersDao) Get(userID uint64) (entity.User, error) {
	user := entity.User{}
	err := u.instance.GetClient().Where("status = ? AND user_id = ?", entity.UserStatusActive, userID).First(&user).Error
	return user, err
}

// GetByIDAndEmail is get by id and email
func (u *UsersDao) GetByIDAndEmail(userID uint64, email string) (entity.User, error) {
	user := entity.User{}
	err := u.instance.GetClient().Where("status = ? AND user_id = ? AND email = ?", entity.UserStatusActive, userID, email).First(&user).Error
	return user, err
}

// GetByEmail is get by  email
func (u *UsersDao) GetByEmail(email string) (entity.User, error) {
	user := entity.User{}
	err := u.instance.GetClient().Where("status = ? AND email = ?", entity.UserStatusActive, email).First(&user).Error
	return user, err
}

// Create is create new user
func (u *UsersDao) Create(name, email, password string) error {
	user, err := u.set(name, email, password)
	if err != nil {
		return err
	}
	tx := u.instance.GetClient().Begin()
	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// Update is update exist user
func (u *UsersDao) Update(userID uint64, name, email, password string) error {
	user, err := u.set(name, email, password)
	if err != nil {
		return err
	}
	return u.instance.GetClient().Model(&entity.User{}).Where(
		"status = ? AND user_id = ?",
		entity.UserStatusActive,
		userID,
	).Updates(user).Error

}

// InActive is update exist user
func (u *UsersDao) InActive(userID uint64) error {
	return u.instance.GetClient().Model(&entity.User{}).Where(
		"status = ? AND user_id = ?",
		entity.UserStatusActive,
		userID,
	).Update("status", entity.UserStatusInActive).Error

}

// Auth is authorize user
func (u *UsersDao) Auth(email, password string) (entity.User, error) {
	user := entity.User{}
	if err := u.instance.GetClient().Where("status = ? AND email = ?", entity.UserStatusActive, email).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	if err := u.comparePassword(user, password); err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// ComparePassword compares input password to roomdata password
func (u *UsersDao) comparePassword(user entity.User, password string) error {
	return actor.PasswordOperator{}.ComparePassword(user.Password, password, user.Salt)
}

func (u *UsersDao) set(name, email, password string) (entity.User, error) {
	hashedpassword, salt, err := actor.PasswordOperator{}.GetHashedPassword(password)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		Name:     name,
		Email:    email,
		Password: hashedpassword,
		Salt:     salt,
	}, nil
}

// RecordNotFoundError is check error as record not found
func (u *UsersDao) RecordNotFoundError(err error) bool {
	return u.instance.RecordNotFoundError(err)
}
