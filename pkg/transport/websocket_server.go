package transport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
	"github.com/WayneShenHH/toolsgo/pkg/transport/middleware"
	"github.com/WayneShenHH/toolsgo/pkg/transport/router"
)

// WebsocketServer websocket server
type WebsocketServer interface {
	Run()
}

// websocketServer websocket server
type websocketServer struct {
	router router.Router
	engine *gin.Engine
	cfg    *environment.WebsocketConfig
}

// NewWebsocketServer websocket server 建構子
func NewWebsocketServer(router router.Router, cfg *environment.WebsocketConfig) WebsocketServer {
	engine := gin.New()
	gin.SetMode(gin.ReleaseMode)
	ctx := context.Background()
	logger.Debug("load middlewares !!")

	engine = loadCommon(
		engine,
		middleware.Config(ctx),
		middleware.TimeZone(ctx),
	)

	router.LoadServiceRouter(
		engine,
	)

	return &websocketServer{
		router: router,
		engine: engine,
		cfg:    cfg,
	}
}

// Run start ws server and block process
func (s *websocketServer) Run() {
	started := make(chan bool)

	go func() {
		if err := http.ListenAndServe(s.cfg.Addr, s.engine); err != nil {
			logger.Warning(fmt.Sprintf("transport/Run ListenAndServe %v", err))
		}
	}()

	go func() {
		if err := pingServer(s.cfg.PingAddr); err != nil {
			logger.Fatal("The router has no response, or it might took too long to start up.")
		}
		logger.Info("The router has been deployed successfully.")
		started <- true
	}()

	go func() {
		<-started
		logger.Debug("deploy success")
	}()

	select {}
}
