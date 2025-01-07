package client

import (
	"github.com/howood/moggiecollector/infrastructure/client/datastore"
	"github.com/howood/moggiecollector/library/utils"
)

const RecordNotFoundMsg = "record not found"

// DatastorAssessor struct
type DatastorAssessor struct {
	Instance datastore.DatastoreInstance
}

// NewDatastorAssessor creates a new CacheAssessor
func NewDatastorAssessor() *DatastorAssessor {
	var I *DatastorAssessor
	switch utils.GetOsEnv("DATASTORE_TYPE", "yogabytedb") {
	case "yogabytedb":
		I = &DatastorAssessor{
			Instance: datastore.NewYugaByteDBClient(),
		}
	default:
		I = &DatastorAssessor{
			Instance: datastore.NewYugaByteDBClient(),
		}
	}
	return I
}
