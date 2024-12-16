package client

import (
	"context"

	"github.com/howood/moggiecollector/infrastructure/client/datastore"
	"github.com/howood/moggiecollector/library/utils"
)

const RecordNotFoundMsg = "record not found"

// DatastorAssessor struct
type DatastorAssessor struct {
	Instance datastore.DatastoreInstance
	ctx      context.Context
}

// NewDatastorAssessor creates a new CacheAssessor
func NewDatastorAssessor(ctx context.Context) *DatastorAssessor {
	var I *DatastorAssessor
	switch utils.GetOsEnv("DATASTORE_TYPE", "yogabytedb") {
	case "yogabytedb":
		I = &DatastorAssessor{
			Instance: datastore.NewYugaByteDbClient(),
			ctx:      ctx,
		}
	default:
		I = &DatastorAssessor{
			Instance: datastore.NewYugaByteDbClient(),
			ctx:      ctx,
		}
	}
	return I
}
