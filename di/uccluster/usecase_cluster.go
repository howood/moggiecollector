package uccluster

import (
	"github.com/howood/moggiecollector/application/usecase"
	"github.com/howood/moggiecollector/di/dbcluster"
	"github.com/howood/moggiecollector/di/svcluster"
)

// DataStore interface.
type UsecaseCluster struct {
	AuthUC    *usecase.AuthUsecase
	UserUC    *usecase.UserUsecase
	UserMfaUC *usecase.UserMfaUsecase
}

// NewDatastore returns DataStore interface.
func NewUsecaseCluster(datastore dbcluster.DataStore, serviceCluster *svcluster.ServiceCluster) *UsecaseCluster {
	return &UsecaseCluster{
		AuthUC:    usecase.NewAuthUsecase(datastore),
		UserUC:    usecase.NewUserUsecase(datastore),
		UserMfaUC: usecase.NewUserMfaUsecase(datastore, serviceCluster),
	}
}
