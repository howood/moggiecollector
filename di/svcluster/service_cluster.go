package svcluster

import (
	"github.com/howood/moggiecollector/application/service"
	"github.com/howood/moggiecollector/di/dbcluster"
)

// DataStore interface.
type ServiceCluster struct {
	AuthenticatorSV service.AuthenticatorService
	RequestLogSV    *service.RequestLogService
}

// NewDatastore returns DataStore interface.
func NewServiceCluster(datastore dbcluster.DataStore) *ServiceCluster {
	return &ServiceCluster{
		AuthenticatorSV: service.NewAuthenticatorService(datastore),
		RequestLogSV:    service.NewRequestLogService(datastore),
	}
}
