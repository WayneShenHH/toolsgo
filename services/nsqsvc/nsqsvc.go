package nsqsvc

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/tools"
	nsq "github.com/bitly/go-nsq"
)

const (
	nsqTCP        = "127.0.0.1:4150" //:4150 , producer
	nsqHTTP       = "127.0.0.1:4151" //:4151
	nsqLookupTCP  = "127.0.0.1:4160" //:4160
	nsqLookupHTTP = "127.0.0.1:4161" //:4161 , consumer
	channel       = "ch_1"
	topicAll      = "comparer_msg,comparer_timer,offer_msg,offer_timer,spider_msg,spider_timer,broadcast_operator,broadcast_player"
)

var nsqdProducer map[string]*nsq.Producer
var nsqConsumer map[string]*nsq.Consumer

// NsqConsumeWorker worker for consuming message
func NsqConsumeWorker(topic string) {
	NsqConsume(topic, func(msg []byte) {
		var m models.Message
		json.Unmarshal(msg, &m)
		tools.Log(m, time.Now())
	})
	select {}
}

// NsqProduceMessage for produce message from json file
func NsqProduceMessage(topic string) {
	jsonfile := "match_point"
	data := models.Message{}
	bytes := tools.LoadJSON(jsonfile)
	json.Unmarshal(bytes, &data)
	tools.Log(data, time.Now())
	NsqProduce(topic, data)
}

// NsqProduce produce from data
func NsqProduce(topic string, obj interface{}) error {
	if nsqdProducer == nil {
		nsqdProducer = make(map[string]*nsq.Producer)
	}
	_, hasKey := nsqdProducer[topic]
	if !hasKey {
		config := nsq.NewConfig()
		conn, _ := nsq.NewProducer(nsqTCP, config)
		nsqdProducer[topic] = conn
	}
	body, e := json.Marshal(obj)
	e = nsqdProducer[topic].Publish(topic, body)
	return e
}

// NsqAddTopic create topic
func NsqAddTopic(topics ...string) {
	for _, topic := range topics {
		post(nsqHTTP+"/topic/create?topic="+topic, "")
		post(nsqHTTP+"/channel/create?topic="+topic+"&channel="+channel, "")
	}
}

// NsqConsume consume a topic
func NsqConsume(topic string, task func(msg []byte)) {
	if nsqConsumer == nil {
		nsqConsumer = make(map[string]*nsq.Consumer)
	}
	_, hasKey := nsqConsumer[topic]
	if !hasKey {
		config := nsq.NewConfig()
		q, _ := nsq.NewConsumer(topic, channel, config)
		nsqConsumer[topic] = q
	}
	nsqConsumer[topic].AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		task(message.Body)
		return nil
	}))
	// err := nsqConsumer[topic].ConnectToNSQLookupd(nsqLookupHTTP)
	err := nsqConsumer[topic].ConnectToNSQD(nsqTCP)
	if err != nil {
		tools.Log(err.Error())
	}
}
func post(url string, obj interface{}) *http.Response {
	body, _ := json.Marshal(obj)
	reader := strings.NewReader(string(body))
	request, _ := http.NewRequest("POST", "http://"+url, reader)
	client := &http.Client{}
	rsp, _ := client.Do(request)
	return rsp
}
