package middleware

import (
	"strconv"

	"github.com/WayneShenHH/toolsgo/tools/timeutil"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

/*TimeZone frontend add Local-UTC header every request
parse time zone convert to float (some time zone has .5)
and add to gin context
*/
func TimeZone(cli *cli.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		tzString := c.Request.Header.Get("Local-UTC")
		tzUTC, err := strconv.ParseFloat(tzString, 64)
		if err != nil {
			// if parse error , default UTC+8
			tzUTC = 8
		}

		timeutil.ToContext(c, timeutil.TimeZone{
			UTC: tzUTC,
		})
		c.Next()
	}
}
