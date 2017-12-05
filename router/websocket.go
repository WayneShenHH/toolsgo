package router

import (
	"github.com/WayneShenHH/toolsgo/router/middleware"
	"github.com/WayneShenHH/toolsgo/wschannel"

	"github.com/gin-gonic/gin"
)

func loadWebsocket(g *gin.Engine) {
	wschannel.Manager.RegisterClient()
	wschannel.Manager.BroadcastSubscribe(wschannel.Operator)
	wschannel.Manager.BroadcastSubscribe(wschannel.Player)
	p := g.Group("/player")
	p.Use(middleware.AuthMiddleware(middleware.Player))
	{
		p.GET("/ch", func(c *gin.Context) {
			wschannel.WSHandler(wschannel.Player, c.Writer, c.Request)
		})
	}
	op := g.Group("/operator")
	op.Use(middleware.AuthMiddleware(middleware.Operator))
	{
		op.GET("/ch", func(c *gin.Context) {
			wschannel.WSHandler(wschannel.Operator, c.Writer, c.Request)
		})
	}
}
