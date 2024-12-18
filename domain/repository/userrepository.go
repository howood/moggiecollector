package repository

import (
	"context"

	"github.com/howood/moggiecollector/domain/entity"
)

// UserRepository interface
type UserRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetAllWithInActive(ctx context.Context) ([]entity.User, error)
	Get(ctx context.Context, userID uint64) (entity.User, error)
	GetByIDAndEmail(ctx context.Context, userID uint64, email string) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, name, email, password string) error
	Update(ctx context.Context, userID uint64, name, email, password string) error
	InActive(ctx context.Context, userID uint64) error
	Auth(ctx context.Context, email, password string) (entity.User, error)
	RecordNotFoundError(err error) bool
}
