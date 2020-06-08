package router

import (
	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/transport/middleware"
)

func (r *Router) loadWebsocket(g *gin.Engine) {
	go r.traderOddsWS.Start()

	base := g.Group(environment.Setting.HTTP.BaseURL)

	op := base.Group("/trader")
	op.Use(middleware.AuthMiddleware(middleware.Operator))
	{
		op.GET("/sample", gin.WrapF(r.traderOddsWS.WSHandler))
	}
}
