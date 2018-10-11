package router

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/errno"
	"github.com/WayneShenHH/toolsgo/module/sd"
	"github.com/WayneShenHH/toolsgo/router/middleware/header"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g = loadCommon(g, mw...)
	loadGeneralRouter(g)
	loadPlayerRouter(g)
	return g
}

// LoadWS 獨立載入 websocket 因為要從不同 container 啟動
func LoadWS(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g = loadCommon(g, mw...)
	loadWebsocket(g)
	return g
}

func loadCommon(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
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

	// The health check handlers
	// for the service discovery.
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		// svcd.GET("/disk", sd.DiskCheck)
		// svcd.GET("/cpu", sd.CPUCheck)
		// svcd.GET("/ram", sd.RAMCheck)
	}
	// 如果 root 路徑沒有回應，k8s 認定服務不正常
	g.GET("/", sd.HealthCheck)
	if app.Setting.Swagger.Enable {
		g.StaticFile(app.Setting.HTTP.BaseURL+"/swagger.json", app.Setting.Swagger.FilePath)
	}
	return g
}
