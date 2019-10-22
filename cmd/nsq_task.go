package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/WayneShenHH/toolsgo/models"
	"github.com/WayneShenHH/toolsgo/module/mq/nsqsvc"
	"github.com/WayneShenHH/toolsgo/tools"

	"github.com/spf13/cobra"
)

var nsqCmd = &cobra.Command{
	Short: "nsq cmd",
	Long:  `nsq worker`,
	Use:   "nsq",
	Run: func(cmd *cobra.Command, args []string) {
		mq := nsqsvc.New()
		topic := args[0]
		channel := args[1]
		mq.ConsumeWorker(topic, channel, func(data []byte) error {
			fmt.Println(string(data))
			return nil
		})
	},
}
var nsqProduceCmd = &cobra.Command{
	Short: "nsq add message cmd",
	Long:  `nsq add message`,
	Use:   "nsq:msg",
	Run: func(cmd *cobra.Command, args []string) {
		mq := nsqsvc.New()
		topic := args[0]
		jsonfile := args[1]
		bytes := tools.LoadJSON(jsonfile)

		msg := models.Message{}
		json.Unmarshal(bytes, &msg)
		mq.Produce(topic, msg)
	},
}
var nsqAddTopicCmd = &cobra.Command{
	Short: "nsq add topic cmd",
	Long:  `nsq add topic`,
	Use:   "topic:add",
	Run: func(cmd *cobra.Command, args []string) {
		topics := strings.Split(args[0], ",")
		nsqsvc.NsqAddTopic(topics...)
	},
}

func init() {
	RootCmd.AddCommand(nsqCmd)
	RootCmd.AddCommand(nsqAddTopicCmd)
	RootCmd.AddCommand(nsqProduceCmd)
}
