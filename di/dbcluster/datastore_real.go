package dbcluster

import (
	"context"
	"errors"

	"github.com/howood/moggiecollector/domain/model"
	"github.com/howood/moggiecollector/infrastructure/client"
	"github.com/howood/moggiecollector/infrastructure/client/datastore"
	"github.com/howood/moggiecollector/infrastructure/client/datastore/dao"
	"gorm.io/gorm"
)

//nolint:gochecknoglobals
var RecordNotFoundMsg = client.RecordNotFoundMsg

// DataStore interface.
type DataStoreReal struct {
	dsRepository *DataStoreRepository
	dbInstance   datastore.DatastoreInstance
}

// NewDatastore returns DataStore interface.
//
//nolint:ireturn
func NewDatastore() DataStore {
	dataaccessor := client.NewDatastorAssessor()
	tables := []interface{}{
		&model.User{},
		&model.UserMfa{},
		&model.RequestLog{},
	}
	dataaccessor.Instance.Migrate(tables)

	return &DataStoreReal{
		dsRepository: &DataStoreRepository{
			UserRepository:       dao.NewUserDao(),
			UserMfaRepository:    dao.NewUserMfaDao(),
			RequestLogRepository: dao.NewRequestLogDao(),
		},
		dbInstance: dataaccessor.Instance,
	}
}

// DSRepository returns dsRepository.
func (d *DataStoreReal) DSRepository() *DataStoreRepository {
	return d.dsRepository
}

// DBInstanceClient returns DBInstance Client.
func (d *DataStoreReal) DBInstanceClient(ctx context.Context) *gorm.DB {
	return d.dbInstance.GetClient().WithContext(ctx)
}

// RecordNotFoundError is check error as record not found
func (d *DataStoreReal) RecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
