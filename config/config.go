package config

import (
	"fmt"
	"os"

	"github.com/howood/moggiecollector/library/utils"
)

const (
	dbPort int = 5432
)

func DatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOSTNAME"),
		utils.GetOsEnvInt("DB_PORT", dbPort),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PASSWORD"),
	)
}
