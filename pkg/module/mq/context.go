package mq

import (
	"context"
)

type key string

const mqKey key = "mq"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Store associated with this context.
func FromContext(c context.Context) MessageQueueService {
	mq := c.Value(string(mqKey))
	if mq == nil {
		mq = c.Value(mqKey)
	}
	return mq.(MessageQueueService)
}

// ToContext adds the Store to this context if it supports the Setter interface.
func ToContext(c Setter, mq MessageQueueService) {
	c.Set(string(mqKey), mq)
}

// SetContext add store to context
func SetContext(c context.Context, store MessageQueueService) context.Context {
	return context.WithValue(c, mqKey, store)
}
