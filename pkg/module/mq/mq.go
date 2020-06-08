// Package mq message queue client
package mq

// MessageQueueService msg queue service
type MessageQueueService interface {
	Consume(topic, channel string, task func(msg []byte) error) error
	Produce(topic string, obj interface{}) error
	ProduceByte(topic string, obj interface{}) error
	ConsumeWorker(topic, ch string, task func(msg []byte) error)
}
