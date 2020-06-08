package websocket

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	json "github.com/json-iterator/go"

	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
	"github.com/WayneShenHH/toolsgo/pkg/module/mq"
)

// Hub operator odds hub
type Hub interface {
	Start()
	WSHandler(res http.ResponseWriter, req *http.Request)
}

// SocketHub all socket channels
type socketHub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	rooms      map[string]map[*Client]bool
	mq         mq.MessageQueueService
	sync.Mutex
	topic             string
	consumeChannel    string
	getSubscribeRooms func(msg []byte) []string
	getBroadcastRooms func(msg []byte) []string
}

// NewHub ctor operator odds hub
func NewHub(mq mq.MessageQueueService, topic, consumeChannel string, getSubscribeRooms func(msg []byte) []string, getBroadcastRooms func(msg []byte) []string) Hub {
	return &socketHub{
		clients:           make(map[*Client]bool),
		broadcast:         make(chan []byte, 5),
		register:          make(chan *Client),
		unregister:        make(chan *Client),
		rooms:             make(map[string]map[*Client]bool),
		mq:                mq,
		topic:             topic,
		consumeChannel:    consumeChannel,
		getSubscribeRooms: getSubscribeRooms,
		getBroadcastRooms: getBroadcastRooms,
	}
}

// Start start socket hub
func (hub *socketHub) Start() {
	defer func() {
		if err := recover(); err != nil {
			logger.Fatal(fmt.Sprintf("[WS] Socket: offer occurs panic, err : %v", err))
		}
	}()
	go hub.subscribe()

	for {
		select {
		case conn := <-hub.register:
			hub.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "welcome operator"})
			conn.Send <- jsonMessage
		case conn := <-hub.unregister:
			if _, ok := hub.clients[conn]; ok {
				conn.disconnect()
			}

		// 收到訊息的時候先解析是哪一場 match 資料更新，針對特定賽事 room 做推送
		case message := <-hub.broadcast:
			rooms := hub.getBroadcastRooms(message)
			for i := range rooms {
				hub.broadcastToRoom(rooms[i], message)
			}
		}
	}
}

// WSHandler called from gin route
func (hub *socketHub) WSHandler(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if error != nil {
		http.NotFound(res, req)
		return
	}
	client := &Client{
		ID:     uuid.New().String(),
		Socket: conn,
		Send:   make(chan []byte, 256),
		hub:    hub,
	}
	hub.register <- client

	go client.read(hub.getSubscribeRooms)
	go client.write()
	go client.ping()
}

func (hub *socketHub) broadcastToRoom(roomName string, message []byte) {
	for conn := range hub.rooms[roomName] {
		conn.SendMessage(message)
	}
}

// subscribe 訂閱 redis worker/broadcast/message
// 如果有收到賠率變動發送推播 & 如果有收到部位重計發送推播
func (hub *socketHub) subscribe() {
	hub.mq.ConsumeWorker(hub.topic, hub.consumeChannel, func(data []byte) error {
		hub.broadcast <- data
		return nil
	})
}
