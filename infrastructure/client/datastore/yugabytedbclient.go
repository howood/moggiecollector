package datastore

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/howood/moggiecollector/library/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// YugaByteDBClient is YugaByteDb Client
type YugaByteDBClient struct {
	Client *gorm.DB
}

// NewYugaByteDBClient creates a new YugaByteDBClient
func NewYugaByteDBClient() *YugaByteDBClient {
	ret := &YugaByteDBClient{Client: generateConnection()}
	return ret
}

//nolint:mnd
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

//nolint:ireturn
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

func (yc *YugaByteDBClient) GetClient() *gorm.DB {
	return yc.Client
}

// Migrate create initial tables
func (yc *YugaByteDBClient) Migrate(tables []interface{}) {
	for _, tabele := range tables {
		if err := yc.Client.AutoMigrate(tabele); err != nil {
			panic(err)
		}
	}
}
