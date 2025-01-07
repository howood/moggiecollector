package datastore

import "gorm.io/gorm"

// DatastoreInstance interface
//
//nolint:revive
type DatastoreInstance interface {
	GetClient() *gorm.DB
	Migrate(tables []interface{})
}
