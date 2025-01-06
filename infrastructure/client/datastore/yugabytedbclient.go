package datastore

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

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

	dbInstance, err := gorm.Open(postgres.Open(dbURI), gormConfig())
	if err != nil {
		panic(err)
	}
	err = dbInstance.Exec("SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL SERIALIZABLE").Error
	if err != nil {
		panic(err)
	}

	return dbInstance
}

func gormConfig() gorm.Option {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)

	return &gorm.Config{
		Logger: newLogger,
	}
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
