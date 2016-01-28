package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/takeanote/takeanote-api/config"
)

type DbConnectionString struct {
	config           config.Config
	connectionString string
}

var databaseConfigs = []DbConnectionString{
	DbConnectionString{
		config: config.Config{
			DatabaseDriver:   "mysql",
			DatabaseHost:     "172.16.0.1",
			DatabasePort:     "3301",
			DatabaseName:     "mysqldb",
			DatabaseUser:     "mysqluser",
			DatabasePassword: "password1",
		},
		connectionString: "mysqluser:password1@tcp(172.16.0.1:3301)/mysqldb",
	},
	DbConnectionString{
		config: config.Config{
			DatabaseDriver:   "postgres",
			DatabaseHost:     "172.16.0.2",
			DatabasePort:     "3302",
			DatabaseName:     "postgresdb",
			DatabaseUser:     "postgresuser",
			DatabasePassword: "password2",
		},
		connectionString: "host=172.16.0.2 port=3302 user=postgresuser password=password2 dbname=postgresdb sslmode=disable",
	},
	DbConnectionString{
		config: config.Config{
			DatabaseDriver:   "sqlite",
			DatabaseHost:     "172.16.0.3",
			DatabasePort:     "3303",
			DatabaseName:     "sqlitedb.db",
			DatabaseUser:     "sqliteuser",
			DatabasePassword: "password3",
		},
		connectionString: "sqlitedb.db",
	},
	DbConnectionString{
		config: config.Config{
			DatabaseDriver:   "failure",
			DatabaseHost:     "172.16.0.3",
			DatabasePort:     "3303",
			DatabaseName:     "sqlitedb.db",
			DatabaseUser:     "sqliteuser",
			DatabasePassword: "password3",
		},
		connectionString: "sqlitedb.db",
	},
}

func TestGenerateDBConnectionString(t *testing.T) {
	for _, dbTest := range databaseConfigs {
		connectionString, err := generateDBConnectionString(&dbTest.config)
		if dbTest.config.DatabaseDriver == "failure" {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, dbTest.connectionString, connectionString)
		}
	}
}
