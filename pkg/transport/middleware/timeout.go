package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	key = "gin.errhandle"
)

// ErrHandle local timezone info
type ErrHandle struct {
	Timeout chan (bool)
}

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext 從 context 中取得身份驗證資訊
// func FromContext(c context.Context) (ErrHandle, bool) {
// 	t, ok := c.Value(key).(ErrHandle)
// 	return t, ok
// }

// ToContext 將資料傳入 context
func ToContext(c Setter, data ErrHandle) {
	c.Set(key, data)
}

// TimeoutFilter set timeout
func TimeoutFilter(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		finish := make(chan struct{})
		ginerr := ErrHandle{Timeout: make(chan bool, t)}
		go func() {
			defer func() {
				if err := recover(); err != nil {
					c.JSON(500, "error")
					c.Abort()
					finish <- struct{}{}
				}
			}()
			c.Next()
			finish <- struct{}{}
		}()

		select {
		case <-time.After(t):
			ginerr.Timeout <- true
			c.JSON(504, "timeout")
			c.Abort()
		case <-finish:
		}
	}
}
