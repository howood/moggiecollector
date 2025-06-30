package client

import (
	"github.com/howood/moggiecollector/infrastructure/client/datastore"
	"github.com/howood/moggiecollector/library/utils"
)

const RecordNotFoundMsg = "record not found"

// DatastorAssessor struct.
type DatastorAssessor struct {
	Instance datastore.DatastoreInstance
}

// NewDataStoreAccesser creates a new CacheAssessor.
func NewDataStoreAccesser() *DatastorAssessor {
	var I *DatastorAssessor
	switch utils.GetOsEnv("DATASTORE_TYPE", "postgres") {
	case "postgres":
		I = &DatastorAssessor{
			Instance: datastore.NewPostgresClient(),
		}
	default:
		I = &DatastorAssessor{
			Instance: datastore.NewPostgresClient(),
		}
	}
	return I
}
