package nsqsvc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/WayneShenHH/toolsgo/module/mq"
	nsq "github.com/nsqio/go-nsq"
)

var (
	nsqConfig = &app.NsqConfig{
		NsqdTCP:       "127.0.0.1:4150", //:4150 , producer
		NsqdHTTP:      "127.0.0.1:4151",
		NsqLookupTCP:  "127.0.0.1:4160",
		NsqLookupHTTP: "127.0.0.1:4161", //:4161 , consumer
		MaxInFlight:   1,
		Concurrency:   1,
	}
)

type nsqService struct {
	producer *nsq.Producer
}

// New instance nsq client
func New() mq.MessageQueueService {
	config := nsq.NewConfig()
	if app.Setting.Nsq != nil {
		nsqConfig = app.Setting.Nsq
	}
	p, _ := nsq.NewProducer(nsqConfig.NsqdTCP, config)
	return &nsqService{
		producer: p,
	}
}

// Produce from data
func (svc *nsqService) Produce(topic string, obj interface{}) error {
	body, e := json.Marshal(obj)
	e = svc.producer.Publish(topic, body)
	// producer.Stop()
	return e
}

// NsqAddTopic create topic
func NsqAddTopic(topics ...string) {
	for _, topic := range topics {
		post(nsqConfig.NsqdHTTP+"/topic/create?topic="+topic, "")
	}
}

// Consume consume a topic
func (*nsqService) Consume(topic, ch string, task func(msg []byte) error) error {
	if !nsq.IsValidTopicName(topic) {
		return fmt.Errorf("topic doesn't exist")
	}
	config := nsq.NewConfig()
	config.MaxInFlight = nsqConfig.MaxInFlight
	client, err := nsq.NewConsumer(topic, ch, config)
	if err != nil {
		return err
	}
	client.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
		return task(message.Body)
	}), nsqConfig.Concurrency)
	err = client.ConnectToNSQLookupd(nsqConfig.NsqLookupHTTP)
	// err := client.ConnectToNSQD(nsqConfig.NsqdTCP)
	if err != nil {
		return err
	}
	return nil
}

// ConsumeWorker worker
func (svc *nsqService) ConsumeWorker(topic, ch string, task func(msg []byte) error) {
	for {
		e := svc.Consume(topic, ch, task)
		if e == nil {
			break
		}
		logger.Error(e)
		// time.Sleep(time.Duration(10) * time.Second)
	}
	select {}
}
func post(url string, obj interface{}) *http.Response {
	body, _ := json.Marshal(obj)
	reader := strings.NewReader(string(body))
	request, _ := http.NewRequest("POST", "http://"+url, reader)
	client := &http.Client{}
	rsp, _ := client.Do(request)
	return rsp
}
