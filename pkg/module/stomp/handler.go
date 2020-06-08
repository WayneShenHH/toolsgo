package stomp

import (
	"strings"

	sng "github.com/gmallard/stompngo"

	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

// MessageHandler deal with message data from ActiveMQ
type MessageHandler interface {
	Parse(*sng.MessageData)
}

// Skip useless data
func Skip(data string) (skip bool) {
	if strings.Contains(data, `status:accepted`) ||
		strings.Contains(data, `"status":"accepted"`) {
		logger.Info(data)
		skip = true
	} else if strings.Contains(data, `status:rejected`) ||
		strings.Contains(data, `"status":"rejected"`) {
		panic(data)
	} else if data == "" {
		logger.Warning("Receiving a empty message")
		skip = true
	}
	return
}
