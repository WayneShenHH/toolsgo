package router

import (
	"github.com/WayneShenHH/toolsgo/controllers/general/authctrlr"
	"github.com/gin-gonic/gin"
)

func loadGeneralRouter(g *gin.Engine) {
	general := g.Group("/general")
	{
		general.POST("/login", authctrlr.Login)
	}
}
