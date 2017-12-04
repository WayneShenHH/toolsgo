package router

import (
	"github.com/WayneShenHH/toolsgo/router/middleware"
	"github.com/gin-gonic/gin"
)

func loadPlayerRouter(g *gin.Engine) {
	player := g.Group("/player")
	player.Use(middleware.AuthMiddleware(middleware.Player))
	{
	}
}
