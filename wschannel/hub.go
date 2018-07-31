package wschannel

import (
	"encoding/json"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
)

// Manager manage socket connections
var Manager = SocketHub{
	clients:    make(map[WSChannel]map[*Client]bool),
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

// SocketHub all socket channels
type SocketHub struct {
	clients    map[WSChannel]map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// WSChannel ws channel enum
type WSChannel int

const (
	// Operator channel
	Operator WSChannel = iota
	// Player channel
	Player
)

// RegisterClient register
func (manager *SocketHub) RegisterClient() {
	go manager.addClient()
}

// BroadcastSubscribe start socket hub
func (manager *SocketHub) BroadcastSubscribe(channel WSChannel) {
	go subscribe(channel)
	go manager.broadcast(channel)
}
func (manager *SocketHub) addClient() {
	for {
		select {
		case conn := <-manager.Register:
			ch := conn.Channel
			_, hasKey := manager.clients[ch]
			if !hasKey {
				manager.clients[ch] = make(map[*Client]bool)
			}
			manager.clients[ch][conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "welcome "})
			conn.Send <- jsonMessage
		case conn := <-manager.Unregister:
			ch := conn.Channel
			if _, ok := manager.clients[ch][conn]; ok {
				close(conn.Send)
				delete(manager.clients[ch], conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.send(ch, jsonMessage, conn)
			}
		}
	}
}
func (manager *SocketHub) send(channel WSChannel, message []byte, ignore *Client) {
	for conn := range manager.clients[channel] {
		if conn != ignore {
			conn.Send <- message
		}
	}
}
func subscribe(channel WSChannel) {
	keys := app.BuildKeys()
	repo := repositoryimpl.New()
	var key string
	switch channel {
	case Operator:
		key = keys.OperatorChannel.Message
	case Player:
		key = keys.PlayerChannel.Message
	}
	for {
		bytes := repo.Blpop(key)
		Manager.Broadcast <- bytes
	}
}
func (manager *SocketHub) broadcast(channel WSChannel) {
	for {
		select {
		case message := <-manager.Broadcast:
			for conn := range manager.clients[channel] {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.clients[channel], conn)
				}
			}
		}
	}
}
