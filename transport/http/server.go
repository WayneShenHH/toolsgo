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

	"github.com/WayneShenHH/toolsgo/router"
	"github.com/WayneShenHH/toolsgo/router/middleware"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

// ServerFlags 設定相關環境參數
var ServerFlags = []cli.Flag{
	// Common flags.
	cli.StringFlag{
		EnvVar: "KITSVC_NAME",
		Name:   "name",
		Usage:  "the name of the service, exposed for service discovery.",
		Value:  "Service",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_URL",
		Name:   "url",
		Usage:  "the url of the service.",
		Value:  "http://127.0.0.1:18086",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_ADDR",
		Name:   "addr",
		Usage:  "the address of the service (with the port).",
		Value:  "127.0.0.1:18086",
	},
	cli.IntFlag{
		EnvVar: "KITSVC_PORT",
		Name:   "port",
		Usage:  "the port of the service.",
		Value:  18086,
	},
	cli.StringFlag{
		EnvVar: "KITSVC_USAGE",
		Name:   "usage",
		Usage:  "the usage of the service, exposed for service discovery.",
		Value:  "Operations about the users.",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_JWT_SECRET",
		Name:   "jwt-secret",
		Usage:  "the secert used to encode the json web token.",
		Value:  "4Rtg8BPKwixXy2ktDPxoMMAhRzmo9mmuZjvKONGPZZQSaJWNLijxR42qRgq0iBb5",
	},
	cli.IntFlag{
		EnvVar: "KITSVC_MAX_PING_COUNT",
		Name:   "max-ping-count",
		Usage:  "the amount to ping the server before we give up.",
		Value:  20,
	},
	cli.StringFlag{
		EnvVar: "KITSVC_DEBUG",
		Name:   "debug",
		Usage:  "enable the debug mode.",
	},

	// Database flags.
	cli.StringFlag{
		EnvVar: "KITSVC_DATABASE_DRIVER",
		Name:   "database-driver",
		Usage:  "the driver of the database.",
		Value:  "mysql",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_DATABASE_NAME",
		Name:   "database-name",
		Usage:  "the name of the database.",
		Value:  "score_go_develop",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_DATABASE_HOST",
		Name:   "database-host",
		Usage:  "the host of the database (with the port).",
		Value:  "127.0.0.1:3306",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_DATABASE_USER",
		Name:   "database-user",
		Usage:  "the user of the database.",
		Value:  "root",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_DATABASE_PASSWORD",
		Name:   "database-password",
		Usage:  "the password of the database.",
		Value:  "123456",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_DATABASE_CONFIG",
		Name:   "database-config",
		Usage:  "database connection string.",
		Value:  "123456",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_DATABASE_CHARSET",
		Name:   "database-charset",
		Usage:  "the charset of the database.",
		Value:  "utf8",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_DATABASE_LOC",
		Name:   "database-loc",
		Usage:  "the timezone of the database.",
		Value:  "Local",
	},
	cli.BoolFlag{
		EnvVar: "KITSVC_DATABASE_PARSE_TIME",
		Name:   "database-parse_time",
		Usage:  "parse the time.",
	},

	// NSQ flags.
	cli.StringFlag{
		EnvVar: "KITSVC_NSQ_PRODUCER",
		Name:   "nsq-producer",
		Usage:  "the address of the TCP NSQ producer (with the port).",
		Value:  "127.0.0.1:4150",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_NSQ_PRODUCER_HTTP",
		Name:   "nsq-producer-http",
		Usage:  "the address of the HTTP NSQ producer (with the port).",
		Value:  "127.0.0.1:4151",
	},
	cli.StringSliceFlag{
		EnvVar: "KITSVC_NSQ_LOOKUPDS",
		Name:   "nsq-lookupds",
		Usage:  "the address of the NSQ lookupds (with the port).",
		Value: &cli.StringSlice{
			"127.0.0.1:4161",
		},
	},

	// Event store flags.
	cli.StringFlag{
		EnvVar: "KITSVC_ES_SERVER_URL",
		Name:   "es-url",
		Usage:  "the url of the event store server.",
		Value:  "http://127.0.0.1:2113",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_ES_USERNAME",
		Name:   "es-username",
		Usage:  "the username of the event store.",
		Value:  "admin",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_ES_PASSWORD",
		Name:   "es-password",
		Usage:  "the password of the event store.",
		Value:  "changeit",
	},

	// Prometheus flags.
	cli.StringFlag{
		EnvVar: "KITSVC_PROMETHEUS_NAMESPACE",
		Name:   "prometheus-namespace",
		Usage:  "the prometheus namespace.",
		Value:  "service",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_PROMETHEUS_SUBSYSTEM",
		Name:   "prometheus-subsystem",
		Usage:  "the subsystem of the promethues.",
		Value:  "user",
	},

	// Consul flags.
	cli.StringFlag{
		EnvVar: "KITSVC_CONSUL_CHECK_INTERVAL",
		Name:   "consul-check_interval",
		Usage:  "the interval of consul health check.",
		Value:  "30s",
	},
	cli.StringFlag{
		EnvVar: "KITSVC_CONSUL_CHECK_TIMEOUT",
		Name:   "consul-check_timeout",
		Usage:  "the timeout of consul health check.",
		Value:  "1s",
	},
	cli.StringSliceFlag{
		EnvVar: "KITSVC_CONSUL_TAGS",
		Name:   "consul-tags",
		Usage:  "the service tags for consul.",
		Value: &cli.StringSlice{
			"user",
			"micro",
		},
	},
	cli.StringFlag{
		EnvVar: "TEST_TAGS",
		Name:   "test.v",
		Usage:  "the service tags for consul.",
		Value:  "",
	},
	cli.StringFlag{
		EnvVar: "TEST_TIMEOUT_TAGS",
		Name:   "test.timeout",
		Usage:  "test timeout",
		Value:  "",
	},
	cli.StringFlag{
		EnvVar: "TEST_RUN_TARGET",
		Name:   "test.run",
		Usage:  "test run target",
		Value:  "",
	},
}

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

	// Create the Gin engine.
	g := gin.New()
	// Event handlers.
	// event := eventutil.New(g)
	// Websocket handlers.
	// ws := wsutil.New(g)
	// Message queue handlers.
	// mq := mqutil.New(g)

	// Routes.
	fmt.Println("load middlewares !!")

	router.Load(
		// Cores.
		g,
		// event, ws, mq,
		// Middlwares.
		middleware.Store(ctx),
		// middleware.TimeZone(ctx),
		// middleware.Logging(),
		// middleware.Event(ctx, event, replayed, deployed),
		// middleware.MQ(ctx, mq, deployed),
		middleware.Metrics(),
	)
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(ctx); err != nil {

		}
		// Close the `deployed` channel to make it non-blocking.
		started <- true
		close(deployed)
	}()

	// Start to listening the incoming requests.
	return http.ListenAndServe(ctx.String("addr"), g)
	// return nil
}

// pingServer pings the http server to make sure the router is working.
func pingServer(c *cli.Context) error {
	var err error
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		pingURLL := c.String("url") + "/sd/health"
		fmt.Println("pingServer " + pingURLL + " count:" + strconv.Itoa(i))
		resp, err := http.Get(pingURLL)
		if err == nil && resp.StatusCode == 200 {
			fmt.Println("ping server StatusCode:", resp.StatusCode)
			break
		} else {
			err = errors.New("Cannot connect to the router")
		}

		// Sleep for a second to continue the next ping.
		time.Sleep(time.Second)
	}
	return err
}
