// Package websocket hub
package websocket

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	json "github.com/json-iterator/go"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

// Message web socket message
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Client socket connection
type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
	hub    *socketHub
}

func (c *Client) read(getSubscribeRooms func(message []byte) []string) {
	defer func() {
		c.hub.unregister <- c
		if err := c.Socket.Close(); err != nil {
			logger.Warning(fmt.Sprintf("broadcastodds/read defer %v", err))
		}
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			c.hub.unregister <- c
			if errc := c.Socket.Close(); errc != nil {
				logger.Warning(fmt.Sprintf("broadcastodds/read Close %v", errc))
			}
			break
		}

		// 根據 NewHub() 時給予的方法加入特定的 rooms
		// 如果 json unmarshal 正確但是這時候的訊息並不是要訂閱賽事
		// 例如 ping / pong 或是其他會造成 rooms 全部被退出
		// 設定 > 0 才退出的問題是如果畫面沒賽事的時候不會退出先前的 rooms 訂閱
		// 需要修改訊息介面能共用其他狀況來判斷

		rooms := getSubscribeRooms(message)
		for idx := range rooms {
			c.join(rooms[idx])
		}
	}
}

// SendMessage 包裝 client 送出訊息方法，並且加入 recover 避免 closed channel panic
func (c *Client) SendMessage(message []byte) {
	defer func() {
		if err := recover(); err != nil {
			c.disconnect()
		}
	}() // chan被關閉時會出panic
	select {
	case c.Send <- message:
	default:
		c.disconnect()
	}
}

// join room by room name
func (c *Client) join(room string) {
	c.hub.Lock()
	conns := c.hub.rooms[room]
	if conns == nil {
		conns = make(map[*Client]bool)
	}
	conns[c] = true
	c.hub.rooms[room] = conns
	c.hub.Unlock()
}

// leaveAllRooms leave all rooms
func (c *Client) leaveAllRooms() {
	c.hub.Lock()
	for key := range c.hub.rooms {
		delete(c.hub.rooms[key], c)
	}
	c.hub.Unlock()
}

// disconnect 斷開使用者連線，移除相關資料
func (c *Client) disconnect() {
	c.leaveAllRooms()
	delete(c.hub.clients, c)
	defer recoverFromChan() // chan被關閉時會出panic
	close(c.Send)
}

func (c *Client) write() {
	defer func() {
		if err := c.Socket.Close(); err != nil {
			logger.Warning(fmt.Sprintf("broadcastodds/write defer %v", err))
		}
	}()

	for {
		message, ok := <-c.Send
		if !ok {
			if err := c.Socket.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
				logger.Warning(fmt.Sprintf("broadcastodds/write select %v", err))
			}
			return
		}
		if err := c.Socket.WriteMessage(websocket.TextMessage, message); err != nil {
			logger.Warning(fmt.Sprintf("broadcastodds/write WriteMessage %v", err))
		}
	}
}

func (c *Client) ping() {
	defer recoverFromChan() // chan 被關閉時會出 panic
	for {
		time.Sleep(time.Duration(environment.Setting.Websocket.PingDelay) * time.Second)
		msg, _ := json.Marshal(&Message{Content: "ping"})
		c.Send <- msg
	}
}

func recoverFromChan() {
	if err := recover(); err != nil {
		logger.Warning(fmt.Sprintf("broadcastodds/ping recoverFromChan %v", err))
	}
}
