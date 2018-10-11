/*Package http SBODDS odds http api server

This is SBODDS sport match management server.
[Learn about Swagger](http://swagger.wordnik.com) or join the IRC channel '#swagger' on irc.freenode.net.
For this sample, you can use the api key 'special-key' to test the authorization filters
Terms Of Service:
http://cow.bet/terms/
    Schemes:
      http
    Host: 54.199.195.93
    BasePath: /
    Version: 1.0.0
    License:
    Contact:

    Consumes:
    - application/json

    Produces:
    - application/json


swagger:meta
*/
package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/WayneShenHH/toolsgo/router"
	"github.com/WayneShenHH/toolsgo/router/middleware"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

// StartServer init Rest api
func StartServer(ctx *cli.Context, started chan bool) error {
	// `deployed` will be closed when the router is deployed.
	deployed := make(chan bool)
	// // `replayed` will be closed after the events are all replayed.
	// replayed := make(chan bool)

	// Debug mode.
	if !ctx.Bool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	// // Initialize the logger.
	// logger.Init()
	// Create the Gin engine.
	g := gin.New()
	// Event handlers.
	// event := eventutil.New(g)
	// Websocket handlers.
	// ws := wsutil.New(g)
	// Message queue handlers.
	// mq := mqutil.New(g)
	logger.Debug(fmt.Sprintf("enabled sports: %v", app.Setting.HTTP.EnableSports))
	// Routes.
	logger.Debug("load middlewares !!")

	router.Load(
		// Cores.
		g,
		// event, ws, mq,
		// Middlwares.
		middleware.Config(ctx),
		middleware.Store(ctx),
		middleware.TimeZone(ctx),
		// middleware.Logging(),
		// middleware.Event(ctx, event, replayed, deployed),
		// middleware.MQ(ctx, mq, deployed),
		// middleware.Metrics(),
	)
	// Register to the service registry when the events were replayed.
	// go func() {
	// 	<-replayed

	// 	sd.Register(ctx)
	// 	// After the service is registered to the consul,
	// 	// close the `started` channel to make it non-blocking.
	// 	logger.Info("The server has been started.")

	// 	close(started)
	// }()

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(ctx, app.Setting.HTTP.PingAddr); err != nil {
			logger.Fatal("The router has no response, or it might took too long to start up.")
		}
		logger.Info("The router has been deployed successfully.")
		// Close the `deployed` channel to make it non-blocking.
		started <- true
		close(deployed)
	}()

	// Start to listening the incoming requests.
	return http.ListenAndServe(app.Setting.HTTP.Addr, g)
	// return nil
}

// pingServer pings the http server to make sure the router is working.
func pingServer(c *cli.Context, pingAddr string) error {
	var err error
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		pingURLL := pingAddr + "/sd/health"
		logger.Info(fmt.Sprintf("pingServer " + pingURLL + " count:" + strconv.Itoa(i)))
		resp, err := http.Get(pingURLL)
		if err == nil && resp.StatusCode == 200 {
			logger.Info(fmt.Sprintf("ping server StatusCode: %v", resp.StatusCode))
			break
		} else {
			logger.Error("ping return code " + strconv.Itoa(resp.StatusCode))
			err = errors.New("Cannot connect to the router")
		}

		// Sleep for a second to continue the next ping.
		logger.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return err
}

// StartServerWS init Rest api
func StartServerWS(ctx *cli.Context, started chan bool) error {
	// `deployed` will be closed when the router is deployed.
	deployed := make(chan bool)
	// // `replayed` will be closed after the events are all replayed.
	// replayed := make(chan bool)

	// Debug mode.
	if !ctx.Bool("debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	// // Initialize the logger.
	// logger.Init()
	// Create the Gin engine.
	g := gin.New()
	// Event handlers.
	// event := eventutil.New(g)
	// Websocket handlers.
	// ws := wsutil.New(g)
	// Message queue handlers.
	// mq := mqutil.New(g)

	// Routes.
	logger.Debug("load middlewares !!")

	router.LoadWS(
		// Cores.
		g,
		// event, ws, mq,
		// Middlwares.
		middleware.Config(ctx),
		middleware.Store(ctx),
		middleware.TimeZone(ctx),
		// middleware.Logging(),
		// middleware.Event(ctx, event, replayed, deployed),
		// middleware.MQ(ctx, mq, deployed),
		// middleware.Metrics(),
	)
	// Register to the service registry when the events were replayed.
	// go func() {
	// 	<-replayed

	// 	sd.Register(ctx)
	// 	// After the service is registered to the consul,
	// 	// close the `started` channel to make it non-blocking.
	// 	logger.Info("The server has been started.")

	// 	close(started)
	// }()

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(ctx, app.Setting.HTTP.PingWSAddr); err != nil {
			logger.Fatal("The router has no response, or it might took too long to start up.")
		}
		logger.Info("The router has been deployed successfully.")
		// Close the `deployed` channel to make it non-blocking.
		started <- true
		close(deployed)
	}()

	// Start to listening the incoming requests.
	return http.ListenAndServe(app.Setting.HTTP.WSAddr, g)
	// return nil
}
