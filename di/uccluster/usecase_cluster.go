package uccluster

import (
	"github.com/howood/moggiecollector/application/usecase"
	"github.com/howood/moggiecollector/di/dbcluster"
)

// DataStore interface.
type UsecaseCluster struct {
	AuthUC *usecase.AuthUsecase
	UserUC *usecase.UserUsecase
}

// NewDatastore returns DataStore interface.
func NewUsecaseCluster(datastore dbcluster.DataStore) *UsecaseCluster {
	return &UsecaseCluster{
		AuthUC: usecase.NewAuthUsecase(datastore),
		UserUC: usecase.NewUserUsecase(datastore),
	}
}
