package middleware

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/pkg/module/mq"
)

// MessageQueue is a middleware function that initializes the message queue connection to
// the context of every request context.
func MessageQueue(ctx context.Context) gin.HandlerFunc {
	mqconn := mq.FromContext(ctx)
	return func(c *gin.Context) {
		mq.ToContext(c, mqconn)
		c.Next()
	}
}
