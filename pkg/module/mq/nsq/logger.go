package nsq

import (
	nsq "github.com/nsqio/go-nsq"

	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

type nsqLogger struct{}

func (nsqLogger) Output(calldepth int, s string) error {
	switch nsq.LogLevel(calldepth) {
	case nsq.LogLevelDebug:
		logger.Debug(s)
	case nsq.LogLevelInfo:
		logger.Info(s)
	case nsq.LogLevelWarning:
		logger.Warning(s)
	case nsq.LogLevelError:
		logger.Error(s)
	}
	return nil
}
