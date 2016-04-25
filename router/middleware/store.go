package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/takeanote/takeanote-api/config"
	"github.com/takeanote/takeanote-api/store"
	"github.com/takeanote/takeanote-api/store/datastore"
)

// Store is a middleware function that initializes the Datastore and attaches to
// the context of every http.Request.
func Store(cfg *config.PostgreSQL, rdis string) gin.HandlerFunc {
	db := datastore.New(cfg, rdis)

	return func(c *gin.Context) {
		store.ToContext(c, db)

		c.Next()
	}
}
