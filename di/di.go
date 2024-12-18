package di

import (
	"github.com/howood/moggiecollector/domain/entity"
	"github.com/howood/moggiecollector/domain/repository"
	"github.com/howood/moggiecollector/infrastructure/client"
	"github.com/howood/moggiecollector/infrastructure/client/datastore/dao"
)

var RecordNotFoundMsg = client.RecordNotFoundMsg

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
	configureddbstore := DataStore{}
	dataaccessor := client.NewDatastorAssessor()
	tables := []interface{}{
		&entity.User{},
	}
	dataaccessor.Instance.Migrate(tables)
	configureddbstore.User = dao.NewUsersDao(dataaccessor.Instance)

	return configureddbstore, nil
}

// GetDataStore returns DataStore
func GetDataStore() DataStore {
	return dataStore
}
