package datastore

import "gorm.io/gorm"

// DatastoreInstance interface
type DatastoreInstance interface {
	GetClient() *gorm.DB
	Migrate(tables []interface{})
	RecordNotFoundError(err error) bool
}
