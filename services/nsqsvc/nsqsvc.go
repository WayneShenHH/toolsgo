package nsqsvc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	nsq "github.com/nsqio/go-nsq"
)

const (
	nsqTCP        = "127.0.0.1:4150" //:4150 , producer
	nsqHTTP       = "127.0.0.1:4151" //:4151
	nsqLookupTCP  = "127.0.0.1:4160" //:4160
	nsqLookupHTTP = "127.0.0.1:4161" //:4161 , consumer
)

// NsqConsumeWorker worker for consuming message
func NsqConsumeWorker(topic, channel string) {
	NsqConsume(topic, channel, func(msg []byte) error {
		fmt.Println(string(msg))
		return nil
	})
	select {}
}

// NsqProduceMessage for produce message from json file
func NsqProduceMessage(topic, msg string) {
	fmt.Println("Produce a message:", msg, ", send to topic:", topic)
	NsqProduce(topic, msg)
}

// NsqProduce produce from data
func NsqProduce(topic string, obj interface{}) error {
	config := nsq.NewConfig()
	producer, _ := nsq.NewProducer(nsqTCP, config)
	body, e := json.Marshal(obj)
	e = producer.Publish(topic, body)
	producer.Stop()
	return e
}

// NsqAddTopic create topic
func NsqAddTopic(topics ...string) {
	for _, topic := range topics {
		post(nsqHTTP+"/topic/create?topic="+topic, "")
	}
}

//NsqGetTopics get all
func NsqGetTopics() []string {
	bytes := get(nsqLookupHTTP + "/topics")
	r := TopicResponse{}
	json.Unmarshal(bytes, &r)
	fmt.Println(string(bytes))
	return r.Topics
}

// NsqConsume consume a topic
func NsqConsume(topic, ch string, task func(msg []byte) error) {
	list := NsqGetTopics()
	has := false
	for _, v := range list {
		if v == topic {
			has = true
			break
		}
	}
	if !has {
		panic("topic doesn't exist")
	}
	client, _ := nsq.NewConsumer(topic, ch, nsq.NewConfig())
	client.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		return task(message.Body)
	}))
	err := client.ConnectToNSQLookupd(nsqLookupHTTP)
	// err := client.ConnectToNSQD(nsqTCP)
	if err != nil {
		panic(err.Error())
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
func get(url string) []byte {
	rsp, err := http.Get("http://" + url)
	if err != nil {
		panic(err.Error())
	}
	defer rsp.Body.Close()
	body, _ := ioutil.ReadAll(rsp.Body)
	return body
}

// TopicResponse rsp fro topic
type TopicResponse struct {
	Topics []string `json:"topics"`
}
