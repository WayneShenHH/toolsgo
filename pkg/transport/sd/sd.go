// Package sd check http server
package sd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//nolint:unused // 先保留定義
const (
	// B byte
	B = 1
	// KB 1024 byte
	KB = 1024 * B
	// MB 1024 KB
	MB = 1024 * KB
	// GB 1024 MB
	GB = 1024 * MB
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}
