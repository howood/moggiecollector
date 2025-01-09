package dbcluster

import (
	"context"

	"github.com/howood/moggiecollector/domain/repository"
	"gorm.io/gorm"
)

type DataStoreRepository struct {
	UserRepository       repository.UserRepository
	RequestLogRepository repository.RequestLogRepository
}
type DataStore interface {
	DSRepository() *DataStoreRepository
	DBInstanceClient(ctx context.Context) *gorm.DB
	RecordNotFoundError(err error) bool
}
