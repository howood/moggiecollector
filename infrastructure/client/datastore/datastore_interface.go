package datastore

import "github.com/jinzhu/gorm"

// DatastoreInstance interface
type DatastoreInstance interface {
	GetClient() *gorm.DB
	Migrate(tables []interface{})
}
