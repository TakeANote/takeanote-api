package config

import "github.com/kelseyhightower/envconfig"

// Config holds configurations values
type Config struct {
	BaseURL            string `envconfig:"BASE_URL" required:"true"`
	Port               int    `envconfig:"PORT" default:"80"`
	LogLevel           string `envconfig:"LOG_LEVEL" default:"info"`
	Debug              bool   `envconfig:"DEBUG" default:"false"`
	Redis              string `envconfig:"REDIS_ADDR" required:"true"`
	DatabaseDriver     string `envconfig:"DB_DRIVER" required:"true"`
	DatabaseHost       string `envconfig:"DB_HOST"`
	DatabasePort       string `envconfig:"DB_PORT"`
	DatabaseName       string `envconfig:"DB_NAME" required:"true"`
	DatabaseUser       string `envconfig:"DB_USER"`
	DatabasePassword   string `envconfig:"DB_PASSWORD"`
	OAuth2ClientID     string `envconfig:"OAUTH2_CLIENT_ID" required:"true"`
	OAuth2ClientSecret string `envconfig:"OAUTH2_CLIENT_SECRET" required:"true"`
	OAuth2RedirectURI  string `envconfig:"OAUTH2_REDIRECT_URI"`
	OAuth2Refresh      int    `envconfig:"OAUTH2_REFRESH" default:"3600"`
}

// New creates a Config struct from config file or environment variables.
func New() *Config {
	var c Config
	envconfig.MustProcess("tan", &c)
	return &c
}
