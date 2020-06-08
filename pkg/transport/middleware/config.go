package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

// configKey is the key name of the config context in the Gin context.
const configKey = "config"

// Config is a middleware function that initializes the config and attaches to
// the context of every request context.
func Config(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(configKey, ctx)
	}
}
