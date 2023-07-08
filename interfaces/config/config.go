package config

import (
	"context"

	"github.com/howood/moggiecollector/application/actor/datastoreservice"
	"github.com/howood/moggiecollector/application/actor/datastoreservice/dao"
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/repository"
)

var RecordNotFoundMsg = datastoreservice.RecordNotFoundMsg

type DataStore struct {
	User repository.UserRepository
}

var dataStore DataStore

func init() {
	var err error
	dataStore, err = configureDatastore()
	if err != nil {
		panic(err)
	}
}

func configureDatastore() (DataStore, error) {
	ctx := context.Background()
	configureddbstore := DataStore{}
	dataaccessor := datastoreservice.NewDatastorAssessor(ctx)
	tables := []interface{}{
		&entity.User{},
	}
	dataaccessor.Instance.Migrate(tables)
	configureddbstore.User = dao.NewUsersDao(ctx, dataaccessor.Instance)

	return configureddbstore, nil
}

// GetDataStore returns DataStore
func GetDataStore() DataStore {
	return dataStore
}
