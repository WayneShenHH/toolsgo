// Package stomp protocol client
package stomp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	sng "github.com/gmallard/stompngo"
	"github.com/gmallard/stompngo/senv"
	"github.com/google/uuid"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
	"github.com/WayneShenHH/toolsgo/pkg/util/sngheader"
)

// Client handle STOMP connection
type Client struct {
	conn       sng.STOMPConnector
	host       string
	id         string
	ssl        bool
	ctx        context.Context
	cancelFunc context.CancelFunc
	once       sync.Once
}

// NewClient instance Client
func NewClient(cfg environment.Conf) *Client {
	c := &Client{
		id:  uuid.New().String(),
		ssl: cfg.Stomp.SSL,
	}
	c.setCtx()
	return c
}

func (cli *Client) setCtx() {
	cli.ctx, cli.cancelFunc = context.WithCancel(context.Background())
}

// Send send message to broker
func (cli *Client) Send(h sng.Headers, data string) {
	logger.Infof("Cli send '%v'", data)
	err := cli.conn.Send(h, data)
	if err != nil {
		logger.Error(err)
		cli.exitFunc()
	}
}

// Sub subscribe queue or topic
func (cli *Client) Sub(h sng.Headers) <-chan sng.MessageData {
	return cli.subscribe(h)
}

// SubWithHandler subscribe and handle data
func (cli *Client) SubWithHandler(h sng.Headers, hdr ...MessageHandler) {
	ch := cli.subscribe(h)

	go func() {
		for {
			select {
			case <-cli.ctx.Done():
				cli.setCtx()
				return
			case m, ok := <-ch:
				if !ok {
					logger.Error("Message queue closed")
					return
				}
				for idx := range hdr {
					hdr[idx].Parse(&m)
				}
			}
		}
	}()
}

func (cli *Client) subscribe(h sng.Headers) <-chan sng.MessageData {
	logger.Info("Cli start subscribe")
	ch, err := cli.conn.Subscribe(h)
	if err != nil {
		logger.Error(err)
		cli.exitFunc()
	}
	return ch
}

// Unsub stop subscription from broker
func (cli *Client) Unsub(h sng.Headers) error {
	return cli.conn.Unsubscribe(h)
}

// JMSCorrelationID JMSCorrelationID
func (cli *Client) JMSCorrelationID() string { return cli.id }

// Disconnect from broker and close all goroutine which opened by cli
func (cli *Client) Disconnect() {
	logger.Info("Cli disconnect")
	if err := cli.conn.Disconnect(sng.Headers{}); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	cli.cancelFunc()
}

// Connect connect to MQ
func (cli *Client) Connect(host string, connSetting *environment.MQConnect, noSSL bool) (err error) {
	var conn sng.STOMPConnector
	if conn, err = connectTCP(host, connSetting, noSSL); err == nil {
		// go cli.once.Do(func() { cli.monitorSignal() })
		cli.conn = conn
		cli.host = host
	}
	return
}

func connectTCP(host string, connSetting *environment.MQConnect, noSSL bool) (sng.STOMPConnector, error) {
	logger.Infof("Connect to %v", host)
	hap := net.JoinHostPort(host, fmt.Sprint(connSetting.Port)) // join host and port
	conn := tcpConn(hap, noSSL)

	connectHeaders := sngheader.Map(map[string]string{
		sng.HK_LOGIN:          connSetting.Username,
		sng.HK_PASSCODE:       connSetting.Password,
		sng.HK_VHOST:          host,
		sng.HK_ACCEPT_VERSION: senv.Protocol(),
	})

	stompConn, err := sng.NewConnector(conn, connectHeaders)
	if err != nil {
		logger.Errorf("STOMP Connect failed, error:%v", err)
		if stompConn != nil {
			logger.Errorf("Connect Response: %+v", stompConn)
		}
		os.Exit(1)
	}
	return stompConn, err
}

func tcpConn(hap string, noSSL bool) net.Conn {
	n, err := net.DialTimeout(sng.NetProtoTCP, hap, 10*time.Second)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	if noSSL {
		return n
	}

	tc := new(tls.Config)
	tc.InsecureSkipVerify = true
	nc := tls.Client(n, tc)
	if err := nc.Handshake(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	return nc
}

func (cli *Client) monitorSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch)

	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGQUIT:
			logger.Warningf("Get signal: %v", sig)
			cli.exitFunc()
		}
	}
}

func (cli *Client) exitFunc() {
	logger.Warning("Unsubscribe")
	cli.Unsub(sng.Headers{"id", "sub-1"})
	cli.Unsub(sng.Headers{"id", "sub-2"})

	logger.Info("Disconnect")
	cli.Disconnect()
	os.Exit(1)
}
