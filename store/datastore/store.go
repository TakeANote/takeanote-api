package datastore

import (
	"database/sql"
	"time"

	"gopkg.in/redis.v3"

	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Postgres Driver

	"github.com/takeanote/takeanote-api/config"
	"github.com/takeanote/takeanote-api/store"
)

// datastore is an implementation of a store.Store built on top
// of the sql/database driver with a relational database backend.
type datastore struct {
	*gorm.DB
	RedisClient *redis.Client
}

// New creates a database connection for the given datasource configuration
// and returns a new Store.
func New(cfg *config.PostgreSQL, rdis string) store.Store {
	if cfg == nil {
		logrus.Fatal("database connection failed: not configured")
	}
	return From(
		open(cfg),
		rdis,
	)
}

// From returns a Store using an existing database connection.
func From(db *gorm.DB, rdisCfg string) store.Store {
	rdis := redis.NewClient(&redis.Options{
		Addr: rdisCfg,
	})
	return &datastore{db, rdis}
}

// open opens a new database connection with the specified
// driver and connection string and returns a gorm.DB.
func open(cfg *config.PostgreSQL) *gorm.DB {
	db, err := cfg.DB()
	if err != nil {
		logrus.Fatalf("database connection failed: %v", err)
	}

	gorm := setupGorm(db)

	if err := pingDatabase(db); err != nil {
		logrus.Fatalf("database ping attempts failed: %v", err)
	}
	return gorm
}

// pingDatabase is an helper function to ping the database with backoff
// to ensure a connection can be established before we proceed with the
// database setup and migration.
func pingDatabase(db *sql.DB) error {
	var err error
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return nil
		}
		logrus.Infof("database ping failed. retry in 1s")
		time.Sleep(time.Second)
	}
	return err
}

// setupGorm will execute SQL queries to create the app tables.
func setupGorm(db *sql.DB) *gorm.DB {
	gorm, _ := gorm.Open("postgres", db)
	return gorm
}
