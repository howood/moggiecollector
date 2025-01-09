package svcluster

import (
	"github.com/howood/moggiecollector/application/service"
	"github.com/howood/moggiecollector/di/dbcluster"
)

// DataStore interface.
type ServiceCluster struct {
	RequestLogSV *service.RequestLogService
}

// NewDatastore returns DataStore interface.
func NewServiceCluster(datastore dbcluster.DataStore) *ServiceCluster {
	return &ServiceCluster{
		RequestLogSV: service.NewRequestLogService(datastore),
	}
}
