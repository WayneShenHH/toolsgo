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
	NsqTcp        = "127.0.0.1:4150" //:4150 , producer
	NsqHttp       = "127.0.0.1:4151" //:4151
	NsqLookupTcp  = "127.0.0.1:4160" //:4160
	NsqLookupHttp = "127.0.0.1:4161" //:4161 , consumer
	Channel       = "ch_1"
	topic         = "hello_nsq"
)

var nsqdProducer map[string]*nsq.Producer
var nsqConsumer map[string]*nsq.Consumer

func NsqConsumeWorker(topic string) {
	NsqConsume(topic, func(msg []byte) {
		var m models.Message
		json.Unmarshal(msg, &m)
		tools.Log(m)
	})
	select {}
}
func NsqProduceMessage(topic string) {
	jsonfile := "match_point"
	data := models.Message{}
	bytes := tools.LoadJson(jsonfile)
	json.Unmarshal(bytes, &data)
	tools.Log(data, time.Now())
	NsqProduce(topic, data)
}
func NsqProduce(topic string, obj interface{}) error {
	if nsqdProducer == nil {
		nsqdProducer = make(map[string]*nsq.Producer)
	}
	_, hasKey := nsqdProducer[topic]
	if !hasKey {
		config := nsq.NewConfig()
		conn, _ := nsq.NewProducer(NsqTcp, config)
		nsqdProducer[topic] = conn
	}
	body, e := json.Marshal(obj)
	e = nsqdProducer[topic].Publish(topic, body)
	return e
}
func NsqAddTopic(topic string) {
	base := "http://" + NsqHttp + "/"
	post(base+"topic/create?topic="+topic, "")
	post(base+"channel/create?topic="+topic+"&channel="+Channel, "")
}
func NsqConsume(topic string, task func(msg []byte)) {
	if nsqConsumer == nil {
		nsqConsumer = make(map[string]*nsq.Consumer)
	}
	_, hasKey := nsqConsumer[topic]
	if !hasKey {
		config := nsq.NewConfig()
		q, _ := nsq.NewConsumer(topic, Channel, config)
		nsqConsumer[topic] = q
	}
	nsqConsumer[topic].AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		task(message.Body)
		return nil
	}))
	// err := nsqConsumer.ConnectToNSQLookupd(NsqLookupHttp)
	err := nsqConsumer[topic].ConnectToNSQD(NsqTcp)
	if err != nil {
		tools.Log(err.Error())
	}
}
func post(url string, obj interface{}) *http.Response {
	body, _ := json.Marshal(obj)
	reader := strings.NewReader(string(body))
	request, _ := http.NewRequest("POST", url, reader)
	client := &http.Client{}
	rsp, _ := client.Do(request)
	return rsp
}
