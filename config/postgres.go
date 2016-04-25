package config

import (
	"database/sql"
	"fmt"
)

// PostgreSQL holds everything you need to
// connect and interact with a PostgreSQL DB.
type PostgreSQL struct {
	User     string `envconfig:"PGSQL_USER"`
	Password string `envconfig:"PGSQL_PW"`
	Host     string `envconfig:"PGSQL_HOST_NAME"`
	DBName   string `envconfig:"PGSQL_DB_NAME"`
}

var (
	// PgSQLMaxOpenConns will be used to set a PostgreSQL
	// drivers MaxOpenConns value.
	PgSQLMaxOpenConns = 1
	// PgSQLMaxIdleConns will be used to set a PostgreSQL
	// drivers MaxIdleConns value.
	PgSQLMaxIdleConns = 1
)

// DB will attempt to open a sql connection with
// the credentials and the current PgSQLMaxOpenConns
// and PgSQLMaxIdleConns values.
// Users must import a postgres driver in their
// main to use this.
func (m *PostgreSQL) DB() (*sql.DB, error) {
	db, err := sql.Open("postgres", m.String())
	if err != nil {
		return db, err
	}
	db.SetMaxIdleConns(PgSQLMaxIdleConns)
	db.SetMaxOpenConns(PgSQLMaxOpenConns)
	return db, nil
}

// String will return the PostgreSQL connection string.
func (m *PostgreSQL) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		m.User,
		m.Password,
		m.Host,
		m.DBName,
	)
}

// LoadPostgresFromEnv will attempt to load a PostgreSQL object
// from environment variables. If not populated, nil
// is returned.
func LoadPostgresFromEnv() *PostgreSQL {
	var postgres PostgreSQL
	LoadEnvConfig(&postgres)
	if postgres.Host != "" {
		return &postgres
	}
	return nil
}
