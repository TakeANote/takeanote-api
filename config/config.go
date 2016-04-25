package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config is a generic struct to hold information for applications that
// need to connect to databases, log events or emit metrics.
// If you have a use case that does not fit this struct, you can
// make a struct containing just the types that suit your needs and use
// some of the helper functions in this package to load it from the environment.
type Config struct {
	Server *Server

	Redis string `envconfig:"REDIS_ADDR"`

	PostgreSQL *PostgreSQL
}

// EnvAppName is used as a prefix for environment variable
// names when using the LoadXFromEnv funcs.
// It defaults to empty.
var EnvAppName = ""

// LoadConfigFromEnv will attempt to inspect the environment
// of any valid config options and will return a populated
// Config struct with what it found.
// If you need a unique config object and want to use envconfig, you
// will need to run the LoadXXFromEnv for each child struct in
// your config struct. For an example on how to do this, check out the
// guts of this function.
func LoadConfigFromEnv() *Config {
	var app Config
	LoadEnvConfig(&app)
	app.PostgreSQL = LoadPostgresFromEnv()
	app.Server = LoadServerFromEnv()
	return &app
}

// LoadEnvConfig will use envconfig to load the
// given config struct from the environment.
func LoadEnvConfig(c interface{}) {
	err := envconfig.Process(EnvAppName, c)
	if err != nil {
		log.Fatal("unable to load env variable: ", err)
	}
}
