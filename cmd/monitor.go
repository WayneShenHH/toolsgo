package cmd

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/WayneShenHH/toolsgo/services/monitorsvc"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Short: "cmd",
	Long:  `monitor svc on linux,ex:GO_ENV=test /Users/wayneshen/go/bin/waynego  monitor -c "#ch2" -m Alice -i5 -e jj`,
	Use:   "monitor",
	Run: func(cmd *cobra.Command, args []string) {
		repo := repositoryimpl.New()
		svc := monitorsvc.New(repo, machine, emoji, channel, timeInterval)
		fmt.Println(timeInterval, machine, emoji, channel)
		svc.Start()
	},
}
var (
	timeInterval int
	machine      string
	emoji        string
	channel      string
)

func init() {
	monitorCmd.Flags().StringVarP(&machine, "machine", "m", "Linda", "machine name")
	monitorCmd.Flags().StringVarP(&emoji, "emoji", "e", "kiss", "slack emoji name")
	monitorCmd.Flags().StringVarP(&channel, "channel", "c", "#libgo-offer", "slack channel name")
	monitorCmd.Flags().IntVarP(&timeInterval, "interval", "i", 10, "minutes for routine interval")
	RootCmd.AddCommand(monitorCmd)
}
