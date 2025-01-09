package datastore

import (
	"log"
	"os"
	"time"

	"github.com/howood/moggiecollector/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MysqlClient is Mysql Client.
type PostgresClient struct {
	Client *gorm.DB
}

// NewMysqlClient creates a new MysqlClient.
func NewPostgresClient() *PostgresClient {
	ret := &PostgresClient{Client: generateConnection()}

	return ret
}

func generateConnection() *gorm.DB {
	dbInstance, err := gorm.Open(postgres.Open(dsn()), gormConfig())
	if err != nil {
		panic(err)
	}

	if err := dbInstance.Exec("SET SESSION CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL SERIALIZABLE").Error; err != nil {
		panic(err)
	}

	if err := dbInstance.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		panic(err)
	}

	return dbInstance
}

//nolint:ireturn
func gormConfig() gorm.Option {
	var loglevel logger.LogLevel

	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		loglevel = logger.Info
	case "INFO":
		loglevel = logger.Info
	case "WARN":
		loglevel = logger.Warn
	case "ERROR":
		loglevel = logger.Error
	default:
		loglevel = logger.Silent
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  loglevel,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)

	return &gorm.Config{
		Logger: newLogger,
	}
}

func (pg *PostgresClient) GetClient() *gorm.DB {
	return pg.Client
}

func (pg *PostgresClient) Migrate(tables []interface{}) {
	for _, tabele := range tables {
		err := pg.Client.AutoMigrate(tabele)
		if err != nil {
			panic(err)
		}
	}
}

func dsn() string {
	return config.DatabaseDSN()
}
