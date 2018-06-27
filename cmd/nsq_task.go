package cmd

import (
	"strings"

	"github.com/WayneShenHH/toolsgo/services/nsqsvc"
	"github.com/spf13/cobra"
)

var nsqCmd = &cobra.Command{
	Short: "nsq cmd",
	Long:  `nsq worker`,
	Use:   "nsq",
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		channel := args[1]
		nsqsvc.NsqConsumeWorker(topic, channel)
	},
}
var nsqProduceCmd = &cobra.Command{
	Short: "nsq add message cmd",
	Long:  `nsq add message`,
	Use:   "nsq:msg",
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		msg := args[1]
		nsqsvc.NsqProduceMessage(topic, msg)
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
var nsqGetTopicCmd = &cobra.Command{
	Short: "nsq get topic cmd",
	Long:  `nsq get topic`,
	Use:   "topic:all",
	Run: func(cmd *cobra.Command, args []string) {
		nsqsvc.NsqGetTopics()
	},
}

func init() {
	RootCmd.AddCommand(nsqCmd)
	RootCmd.AddCommand(nsqAddTopicCmd)
	RootCmd.AddCommand(nsqProduceCmd)
	RootCmd.AddCommand(nsqGetTopicCmd)
}
