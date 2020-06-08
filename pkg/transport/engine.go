package transport

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/errno"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
	"github.com/WayneShenHH/toolsgo/pkg/transport/middleware/header"
	"github.com/WayneShenHH/toolsgo/pkg/transport/sd"
)

// loadCommon loads common
func loadCommon(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// func Load(g *gin.Engine) *gin.Engine {
	// Middlewares.
	g.Use(header.NoCache)
	g.Use(header.Options)
	g.Use(header.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
		logger.Debug(fmt.Sprintf("The incorrect API route. %s%s", c.Request.Host, c.Request.URL))
		errno.Abort(errno.ErrNotFound, nil, c)
	})
	// default is "debug/pprof"
	pprof.Register(g, "debug/pprof")

	// The health check handlers
	// for the service discovery.
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}
	// 如果 root 路徑沒有回應，k8s 認定服務不正常
	g.GET("/", sd.HealthCheck)
	if config.Swagger().Enable {
		g.StaticFile(config.HTTP().BaseURL+"/swagger.json", config.Swagger().FilePath)
	}
	return g
}
