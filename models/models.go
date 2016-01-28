package models

import (
	"fmt"
	"strings"

	"github.com/takeanote/takeanote-api/config"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql" // MySQL Driver for gorm
	_ "github.com/lib/pq"              // Postgres Driver for gorm
	_ "github.com/mattn/go-sqlite3"    // SQLite Driver for gorm
)

func generateDBConnectionString(config *config.Config) (string, error) {
	switch strings.ToLower(config.DatabaseDriver) {
	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DatabaseHost, config.DatabasePort, config.DatabaseUser,
			config.DatabasePassword, config.DatabaseName), nil
	case "mysql":
		return fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			config.DatabaseUser, config.DatabasePassword, config.DatabaseHost,
			config.DatabasePort, config.DatabaseName), nil
	case "sqlite", "sqlite3":
		return fmt.Sprintf("%s", config.DatabaseName), nil
	}
	return "", fmt.Errorf("unknown DB driver: %s\n", config.DatabaseDriver)
}

// OpenDBWithConfig open a gorm.DB connection thanks to config.Database
func OpenDBWithConfig(config *config.Config) (dbConn gorm.DB, err error) {
	var connectionString string
	if connectionString, err = generateDBConnectionString(config); err != nil {
		return gorm.DB{}, err
	}

	if dbConn, err = gorm.Open(config.DatabaseDriver, connectionString); err != nil {
		return dbConn, err
	}
	return dbConn, err
}
