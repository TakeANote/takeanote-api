package config

// Server holds info required to configure a Swydd server.
type Server struct {
	// GOMAXPROCS can be used to override the default GOMAXPROCS (runtime.NumCPU).
	GOMAXPROCS *int `envconfig:"SERVER_GOMAXPROCS"`
	// HTTPPort is the port the server implementation will serve HTTP over.
	HTTPPort int `envconfig:"HTTP_PORT"`
	// LogLevel will override the default log level of 'info'.
	LogLevel string `envconfig:"LOG_LEVEL"`
}

// LoadServerFromEnv will attempt to load a Server object
// from environment variables. If not populated, nil
// is returned.
func LoadServerFromEnv() *Server {
	var server Server
	LoadEnvConfig(&server)
	if server.HTTPPort != 0 {
		return &server
	}
	return nil
}
