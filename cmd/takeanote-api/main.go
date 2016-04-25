package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"

	"github.com/takeanote/takeanote-api/config"
	"github.com/takeanote/takeanote-api/router"
	"github.com/takeanote/takeanote-api/router/middleware"
)

func main() {
	var cfg *config.Config

	cfg = config.LoadConfigFromEnv()
	if cfg.Server != nil && cfg.Server.GOMAXPROCS != nil {
		runtime.GOMAXPROCS(*cfg.Server.GOMAXPROCS)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	handler := router.Load(
		ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true),
		middleware.Store(cfg.PostgreSQL, cfg.Redis),
	)

	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.HTTPPort), handler))
}
