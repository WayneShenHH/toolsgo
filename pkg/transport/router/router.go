// Package router gin http router interface
package router

import (
	"github.com/gin-gonic/gin"
)

// Router 設定 route 對應 endpoint
type Router interface {
	// LoadServiceRouter 由各自 service 自己定義所需的內部 routes
	LoadServiceRouter(g *gin.Engine)
}
