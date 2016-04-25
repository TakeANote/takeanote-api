package models

import (
	"fmt"
	"strings"

	"github.com/takeanote/takeanote-api/config"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"              // Postgres Driver for gorm
)

func generateDBConnectionString(config *config.Config) (string, error) {
	switch strings.ToLower(config.DatabaseDriver) {
	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DatabaseHost, config.DatabasePort, config.DatabaseUser,
			config.DatabasePassword, config.DatabaseName), nil
	}
	return "", fmt.Errorf("unknown DB driver: %s\n", config.DatabaseDriver)
}

// OpenDBWithConfig open a gorm.DB connection thanks to config.Database
func OpenDBWithConfig(config *config.Config) (dbConn *gorm.DB, err error) {
	var connectionString string
	if connectionString, err = generateDBConnectionString(config); err != nil {
		return nil, err
	}

	if dbConn, err = gorm.Open(config.DatabaseDriver, connectionString); err != nil {
		return nil, err
	}
	return dbConn, err
}
