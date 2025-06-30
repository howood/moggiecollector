package dbcluster

import (
	"context"
	"errors"

	"github.com/howood/moggiecollector/infrastructure/client"
	"github.com/howood/moggiecollector/infrastructure/client/datastore"
	"github.com/howood/moggiecollector/infrastructure/client/datastore/dao"
	"gorm.io/gorm"
)

// DataStore interface.
type DataStoreTest struct {
	dbInstance   datastore.DatastoreInstance
	testTx       *gorm.DB
	dsRepository *DataStoreRepository
}

// NewDatastore returns DataStore interface.
//
//nolint:ireturn
func NewDatastoreForTest(testTx *gorm.DB) DataStore {
	dataaccessor := client.NewDatastorAssessor()

	return &DataStoreTest{
		dsRepository: &DataStoreRepository{
			UserRepository:       dao.NewUserDao(),
			UserMfaRepository:    dao.NewUserMfaDao(),
			RequestLogRepository: dao.NewRequestLogDao(),
		},
		dbInstance: dataaccessor.Instance,
		testTx:     testTx,
	}
}

// DSRepository returns dsRepository.
func (d *DataStoreTest) DSRepository() *DataStoreRepository {
	return d.dsRepository
}

// DBInstanceClient returns textTx.
func (d *DataStoreTest) DBInstanceClient(_ context.Context) *gorm.DB {
	return d.testTx
}

// RecordNotFoundError is check error as record not found
func (d *DataStoreTest) RecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
