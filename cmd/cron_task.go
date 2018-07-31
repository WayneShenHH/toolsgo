package cmd

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/WayneShenHH/toolsgo/services/schedulesvc"
	"github.com/spf13/cobra"
)

var workerScheduleCmd = &cobra.Command{
	Short: "Worker schedule",
	Long:  "Schedured process",
	Use:   "worker:cron",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
		fmt.Println(cmd.Long)
		repo := repositoryimpl.New()
		svc := schedulesvc.New(repo)
		svc.Start()
	},
}
var clearOddsCmd = &cobra.Command{
	Short: "Worker schedule",
	Long:  "Schedured process",
	Use:   "worker:clear",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
		fmt.Println(cmd.Long)
		repo := repositoryimpl.New()
		svc := schedulesvc.New(repo)
		svc.ClearDataTask()
	},
}

func init() {
	RootCmd.AddCommand(workerScheduleCmd)
	RootCmd.AddCommand(clearOddsCmd)
}
