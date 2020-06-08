/*Package transport FSBS api server

This is FSBS sport betting system

Schemes: https, http
Host: s1.b5k.io
BasePath: /api/ums

Security:
- bearer

SecurityDefinitions:
bearer:
  type: apiKey
  name: Authorization
  in: header

swagger:meta
*/
package transport

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
	"github.com/WayneShenHH/toolsgo/pkg/transport/middleware"
	"github.com/WayneShenHH/toolsgo/pkg/transport/router"
)

// APIServer 客端 & 控端 api server
type APIServer struct {
	router  router.Router
	http    *http.Server
	started chan bool
}

// NewAPIServer 客端 & 控端 api server 建構子
func NewAPIServer(router router.Router, config *environment.Setting) *APIServer {
	engine := gin.New()
	ctx := context.Background()
	started := make(chan bool)

	ginSetting := config.Gin
	httpcfg := config.HTTP
	gin.SetMode(gin.ReleaseMode)
	// Create the Gin engine.
	if ginSetting.LogEnable {
		engine.Use(gin.Logger())
	}
	engine.Use(gin.Recovery())

	logger.Debug("load middlewares !!")
	engine = loadCommon(
		engine,
		middleware.Config(ctx),
		middleware.TimeZone(ctx),
	)
	router.LoadServiceRouter(
		engine,
	)

	return &APIServer{
		router: router,
		http: &http.Server{
			Addr:           httpconfig.Addr,
			Handler:        engine,
			ReadTimeout:    time.Duration(httpcfg.TimeoutMS) * time.Millisecond,
			WriteTimeout:   time.Duration(httpcfg.TimeoutMS) * time.Millisecond,
			MaxHeaderBytes: 1 << 20,
		},
		started: started,
	}
}

// Run http server
func (s *APIServer) Run() {
	// tested := make(chan bool)
	// deployed := make(chan bool)

	defer func() {
		if err := recover(); err != nil {
			logger.Fatalf("http server crash：%v", err)
		}
	}()
	go func() {
		if err := s.http.ListenAndServe(); err != nil {
			logger.Fatal("The router has no response, or it might took too long to start up.")
		}
	}()
	// ping server 回應確認對外
	go func() {
		time.Sleep(time.Second)
		if err := pingServer("http://127.0.0.1" + s.http.Addr); err != nil {
			logger.Fatal("The router has no response, or it might took too long to start up.")
		}
		logger.Info("The router has been deployed successfully.")
		// Close the `deployed` channel to make it non-blocking.
		s.started <- true
		// close(deployed)
	}()
	go func() {
		<-s.started
		logger.Debug("deploy success")
		// tested <- true
	}()
	// <-tested
	select {}
}

// pingServer pings the http server to make sure the router is working.
func pingServer(pingAddr string) error {
	var err error

	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		pingURLL := pingAddr + "/sd/health"
		logger.Info(fmt.Sprintf("pingServer %v count:%d", pingURLL, i))
		var res *http.Response
		res, err = http.Get(pingURLL) //nolint
		if err != nil {
			fmt.Println(err)
			// Sleep for a second to continue the next ping.
			logger.Info("waiting for the router, retry in 1 second.")
			time.Sleep(time.Second)
			continue
		}
		defer func() {
			logger.Info(res.Body.Close())
		}()
		if res.StatusCode == 200 {
			logger.Info(fmt.Sprintf("ping server StatusCode: %v", res.StatusCode))
			break
		}
		logger.Error("ping return code " + strconv.Itoa(res.StatusCode))
		err = errors.New("cannot connect to the router")
		logger.Error(err)
	}
	return err
}
