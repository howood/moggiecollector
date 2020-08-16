package datastore

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/howood/moggiecollector/library/utils"
)

// YugaByteDbClient is YugaByteDb Client
type YugaByteDbClient struct {
	Client *gorm.DB
}

// NewYugaByteDbClient creates a new YugaByteDbClient
func NewYugaByteDbClient() *YugaByteDbClient {
	ret := &YugaByteDbClient{Client: generateConnection()}
	return ret
}

func generateConnection() *gorm.DB {
	var err error
	dbURI := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("YUGABYTEDB_HOSTNAME"),
		utils.GetOsEnvInt("YUGABYTEDB_PORT", 5433),
		os.Getenv("YUGABYTEDB_USER"),
		os.Getenv("YUGABYTEDB_DBNAME"),
		os.Getenv("YUGABYTEDB_PASSWORD"),
	)

	dbInstance, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic(err)
	}
	dbInstance.LogMode(true)
	err = dbInstance.Exec("SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL SERIALIZABLE").Error
	if err != nil {
		panic(err)
	}

	return dbInstance
}

func (yc *YugaByteDbClient) GetClient() *gorm.DB {
	return yc.Client
}

// Migrate create initial tables
func (yc *YugaByteDbClient) Migrate(tables []interface{}) {
	for _, tabele := range tables {
		yc.Client.AutoMigrate(tabele)
	}

}
