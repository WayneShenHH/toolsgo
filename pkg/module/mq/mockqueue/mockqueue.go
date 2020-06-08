// Package mockqueue mock message queue
package mockqueue

import (
	"fmt"
	"sync"
	"time"

	json "github.com/json-iterator/go"

	"github.com/google/uuid"
)

var mux *sync.Mutex

// New instance nsq client
func New() *MockNSQD {
	n := &MockNSQD{
		buffer:     make(map[string][][]byte),
		clients:    make(map[string]map[string]*client),
		broadcast:  make(chan []byte, 5),
		register:   make(chan *client),
		unregister: make(chan *client),
		produceCh:  make(chan bool, 1),
		finishChan: make(map[string]chan string),
		history:    make(map[string][]string),
	}
	mux = &sync.Mutex{}
	go n.listen()
	go n.checkBuffer()
	return n
}

func (svc *MockNSQD) listen() {
	go svc.listenRegister()
	go svc.listenUnregister()
}

func (svc *MockNSQD) listenRegister() {
	for {
		conn := <-svc.register
		mux.Lock()
		_, exist := svc.clients[conn.topic]
		if !exist {
			svc.clients[conn.topic] = make(map[string]*client)
		}
		svc.clients[conn.topic][conn.channel] = conn
		fmt.Printf("register: %s, topic: %s\n", conn.ID, conn.topic)
		mux.Unlock()
	}
}
func (svc *MockNSQD) listenUnregister() {
	for {
		conn := <-svc.unregister
		mux.Lock()
		_, exist := svc.clients[conn.topic]
		if exist {
			channels := svc.clients[conn.topic]
			delete(channels, conn.channel)
			svc.clients[conn.topic] = channels
			fmt.Println("unregister:", conn.ID)
		}
		mux.Unlock()
	}
}

func (svc *MockNSQD) checkBuffer() {
	for {
		for topic, msglist := range svc.buffer {
			if conns, exist := svc.clients[topic]; exist {
				mux.Lock()
				for _, msg := range msglist {
					for _, conn := range conns {
						conn.Send <- msg
					}
				}
				svc.buffer[topic] = make([][]byte, 0)
				mux.Unlock()
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

type client struct {
	ID      string
	Send    chan []byte
	topic   string
	channel string
}

// MockNSQD mock queue struct
type MockNSQD struct {
	buffer     map[string][][]byte
	history    map[string][]string // 收到的訊息紀錄 map[topic]
	clients    map[string]map[string]*client
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
	produceCh  chan bool
	finishChan map[string]chan string
}

// Produce from data
func (svc *MockNSQD) Produce(topic string, obj interface{}) error {
	body, e := json.Marshal(obj)

	mux.Lock()
	svc.history[topic] = append(svc.history[topic], string(body))
	queue := svc.buffer[topic]
	queue = append(queue, body)
	svc.buffer[topic] = queue
	mux.Unlock()

	return e
}

// GetHistory 回傳 mock NSQ topic 收到的紀錄
func (svc *MockNSQD) GetHistory(topic string) []string {
	mux.Lock()
	history := svc.history[topic]
	mux.Unlock()
	return history
}

// Flush clear history
func (svc *MockNSQD) Flush() {
	mux.Lock()
	svc.history = make(map[string][]string)
	svc.buffer = make(map[string][][]byte)
	svc.clients = make(map[string]map[string]*client)
	mux.Unlock()
}

// Consume the data, mock-queue cannot register multiple consumer with same topic & channel
// Must unregister after task done
// Block until previous consumer finished
func (svc *MockNSQD) Consume(topic, ch string, task func(msg []byte) error) error {
	for {
		if !svc.block(topic, ch) {
			break
		}
	}
	mux.Lock()
	c := &client{
		ID:      uuid.New().String(),
		topic:   topic,
		channel: ch,
		Send:    make(chan []byte, 10)}
	svc.register <- c
	go c.handler(svc, topic, task)
	mux.Unlock()
	return nil
}
func (svc *MockNSQD) block(topic, ch string) bool {
	mux.Lock()
	defer mux.Unlock()
	ex := false
	if _, ex = svc.clients[topic]; !ex {
		return false
	}
	if _, ex = svc.clients[topic][ch]; !ex {
		return false
	}
	if svc.clients[topic][ch] == nil {
		return false
	}
	return true
}

func (c *client) handler(svc *MockNSQD, topic string, task func(msg []byte) error) {
	msg := <-c.Send
	err := task(msg)
	if err != nil {
		panic(fmt.Errorf("task error"))
	}

	svc.unregister <- c

	if _, exist := svc.finishChan[topic]; exist {
		svc.finishChan[topic] <- ""
		delete(svc.finishChan, topic)
	}
}

// WaitTopicFinish 傳入正確的 topic name 就會回傳一個 channel 來等到 worker 執行結束
func (svc *MockNSQD) WaitTopicFinish(key string) chan string {
	if _, exists := svc.finishChan[key]; exists {
		panic(fmt.Errorf("key %s exists", key))
	}
	svc.finishChan[key] = make(chan string)
	return svc.finishChan[key]
}

// ConsumeWorker will not stop consuming
func (svc *MockNSQD) ConsumeWorker(topic, ch string, task func(msg []byte) error) {
	c := &client{
		ID:      uuid.New().String(),
		topic:   topic,
		channel: ch,
		Send:    make(chan []byte, 10),
	}
	svc.register <- c
	for {
		c.handler(svc, topic, task)
	}
}
