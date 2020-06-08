// Package middleware gin middleware
package middleware

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/pkg/module/timer"
)

/*TimeZone frontend add Local-UTC header every request
parse time zone convert to float (some time zone has .5)
and add to gin context
*/
func TimeZone(ctx context.Context) gin.HandlerFunc {
	tz := timer.FromContext(ctx)
	return func(c *gin.Context) {
		tzString := c.Request.Header.Get("Local-UTC")
		tzUTC, err := strconv.ParseFloat(tzString, 64)
		if err != nil {
			// if parse error , default UTC+8
			tzUTC = 8
		}
		if tz != nil {
			tz.SetUTC(tzUTC)
		}
		timer.ToContext(c, tz)
		c.Next()
	}
}
