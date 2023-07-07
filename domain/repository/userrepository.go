package repository

import (
	"github.com/howood/moggiecollector/domain/entity"
)

// UserRepository interface
type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetAllWithInActive() ([]entity.User, error)
	Get(userID uint64) (entity.User, error)
	GetByIDAndEmail(userID uint64, email string) (entity.User, error)
	Create(name, email, password string) error
	Update(userID uint64, name, email, password string) error
	InActive(userID uint64) error
	Auth(email, password string) (entity.User, error)
}
