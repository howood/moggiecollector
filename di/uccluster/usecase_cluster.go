package uccluster

import (
	"github.com/howood/moggiecollector/application/usecase"
	"github.com/howood/moggiecollector/di/dbcluster"
)

// DataStore interface.
type UsecaseCluster struct {
	AccountUC *usecase.AccountUsecase
	ClientUC  *usecase.ClientUsecase
}

// NewDatastore returns DataStore interface.
func NewUsecaseCluster(datastore dbcluster.DataStore) *UsecaseCluster {
	return &UsecaseCluster{
		AccountUC: usecase.NewAccountUsecase(datastore),
		ClientUC:  usecase.NewClientUsecase(datastore),
	}
}
