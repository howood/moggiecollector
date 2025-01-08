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

// PostgresClient is PostgreSQL Client
type PostgresClient struct {
	Client *gorm.DB
}

// NewPostgresClient creates a new PostgresClient
func NewPostgresClient() *PostgresClient {
	ret := &PostgresClient{Client: generateConnection()}
	return ret
}

//nolint:mnd
func generateConnection() *gorm.DB {
	var err error
	dbURI := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOSTNAME"),
		utils.GetOsEnvInt("DB_PORT", 5433),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PASSWORD"),
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

func (yc *PostgresClient) GetClient() *gorm.DB {
	return yc.Client
}

// Migrate create initial tables
func (yc *PostgresClient) Migrate(tables []interface{}) {
	for _, tabele := range tables {
		if err := yc.Client.AutoMigrate(tabele); err != nil {
			panic(err)
		}
	}
}
