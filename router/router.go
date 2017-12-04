package router

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/errno"
	"github.com/WayneShenHH/toolsgo/module/sd"
	"github.com/WayneShenHH/toolsgo/router/middleware/header"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// func Load(g *gin.Engine) *gin.Engine {
	// Middlewares.
	g.Use(gzip.Gzip(gzip.DefaultCompression))
	g.Use(gin.Recovery())
	g.Use(header.NoCache)
	g.Use(header.Options)
	g.Use(header.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
		errno.Abort(errno.ErrNotFound, nil, c)
	})

	loadGeneralRouter(g)
	loadPlayerRouter(g)
	// loadAgentRouter(g)
	// loadOperatorRouter(g)
	// loadWebsocket(g)

	// The health check handlers
	// for the service discovery.
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		// svcd.GET("/disk", sd.DiskCheck)
		// svcd.GET("/cpu", sd.CPUCheck)
		// svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
