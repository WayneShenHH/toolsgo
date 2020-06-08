// Package nsq message queue implement with nsq
package nsq

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	json "github.com/json-iterator/go"
	nsq "github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/errors"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
	"github.com/WayneShenHH/toolsgo/pkg/module/mq"
)

type nsqService struct {
	producer  *nsq.Producer
	nsqConfig *environment.NSQConfig
}

// New instance nsq client
func New(nsqConfig *environment.NSQConfig) mq.MessageQueueService {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(nsqConfig.NSQDTCP, config)
	p.SetLogger(nsqLogger{}, nsq.LogLevelWarning)
	if err != nil {
		logger.Error(errors.E(err))
	}
	if nsqConfig.NSQDValid {
		// ping 10 次確認 nsqd 存在再啟動
		for i := 0; i < 10; i++ {
			err = p.Ping()
			if err == nil {
				logger.Info(fmt.Sprintf("ping nsqd success"))
				break
			}
			// 十次還抓不到 nsqd 就 panic
			if i == 9 {
				logger.Fatal(err)
			}
			// Sleep for a second to continue the next ping.
			time.Sleep(time.Second)
		}
	}
	return &nsqService{
		producer:  p,
		nsqConfig: nsqConfig,
	}
}

// Produce from data
func (svc *nsqService) Produce(topic string, obj interface{}) error {
	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	logger.Debugf("topic: %v, produce data:\n%v", topic, string(body))
	err = svc.producer.Publish(topic, body)
	// producer.Stop()
	return err
}

// ProduceByte raw byte from data
func (svc *nsqService) ProduceByte(topic string, obj interface{}) error {
	body, err := GetBytes(obj)
	err = svc.producer.Publish(topic, body)
	if err != nil {
		return err
	}
	return nil
}

// AddTopic create topic
func (svc *nsqService) AddTopic(topics ...string) (err error) {

	for _, topic := range topics {
		resp, err := post(svc.nsqConfig.NSQDHTTP+"/topic/create?topic="+topic, "")
		if err != nil {
			return err
		}
		if err = resp.Body.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Consume consume a topic
func (svc *nsqService) Consume(topic, ch string, task func(msg []byte) error) error {
	if !nsq.IsValidTopicName(topic) {
		return fmt.Errorf("topic doesn't exist")
	}
	config := nsq.NewConfig()
	config.MaxInFlight = svc.nsqConfig.MaxInFlight
	client, err := nsq.NewConsumer(topic, ch, config)
	client.SetLogger(nsqLogger{}, nsq.LogLevelWarning)
	if err != nil {
		return err
	}
	client.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
		return task(message.Body)
	}), svc.nsqConfig.Concurrency)
	err = client.ConnectToNSQLookupd(svc.nsqConfig.NSQLookupHTTP)
	// err := client.ConnectToNSQD(nsqConfig.NSQDTCP)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"NSQLookupHTTP": svc.nsqConfig.NSQLookupHTTP,
		}).Errorf(`ConnectToNSQLookupd: %v`, err)
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
		// logger.Error(e)
		// time.Sleep(time.Duration(10) * time.Second)
	}
	select {}
}

func post(url string, obj interface{}) (*http.Response, error) {
	var err error
	body, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(string(body))
	request, err := http.NewRequest("POST", "http://"+url, reader)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	rsp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetBytes produce bytes
func GetBytes(v interface{}) ([]byte, error) {
	str := fmt.Sprint(v)
	b := []byte(str)
	return b, nil
}
