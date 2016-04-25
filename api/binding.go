package api

import (
	"github.com/gin-gonic/gin"
	"github.com/takeanote/takeanote-api/binding"
)

// BindJSON is a shortcut for c.BindWith(obj, binding.JSON)
func BindJSON(c *gin.Context, obj interface{}) error {
	return BindWith(c, obj, binding.JSON)
}

// BindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func BindWith(c *gin.Context, obj interface{}, b binding.Binding) error {
	if err := b.Bind(c.Request, obj); err != nil {
		return err
	}
	return nil
}
