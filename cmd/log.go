package cmd

import (
	"strconv"

	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/WayneShenHH/toolsgo/services/logsvc"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Short: "log cmd",
	Long:  `read log msg`,
	Use:   "log",
	Run: func(cmd *cobra.Command, args []string) {
		mid := args[0]
		i, _ := strconv.Atoi(mid)
		repo := repositoryimpl.New()
		log := logsvc.New(repo)
		log.Read(int(i), int(i))
	},
}

func init() {
	RootCmd.AddCommand(logCmd)
}
