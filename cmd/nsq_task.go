package cmd

import (
	"strings"

	"github.com/WayneShenHH/toolsgo/services"
	"github.com/WayneShenHH/toolsgo/services/nsqsvc"
	"github.com/spf13/cobra"
)

var nsqCmd = &cobra.Command{
	Short: "nsq cmd",
	Long:  `nsq worker`,
	Use:   "nsq",
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		go services.CheckStatus()
		nsqsvc.NsqConsumeWorker(topic)
	},
}
var nsqTopicCmd = &cobra.Command{
	Short: "nsq add topic cmd",
	Long:  `nsq add topic`,
	Use:   "nsq:topic",
	Run: func(cmd *cobra.Command, args []string) {
		topics := strings.Split(args[0], ",")
		nsqsvc.NsqAddTopic(topics...)
	},
}
var nsqProduceCmd = &cobra.Command{
	Short: "nsq add message cmd",
	Long:  `nsq add message`,
	Use:   "nsq:msg",
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		nsqsvc.NsqProduceMessage(topic)
	},
}

func init() {
	RootCmd.AddCommand(nsqCmd)
	RootCmd.AddCommand(nsqTopicCmd)
	RootCmd.AddCommand(nsqProduceCmd)
}
