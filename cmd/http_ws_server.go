package cmd

import (
	"os"

	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/WayneShenHH/toolsgo/transport/http"

	"github.com/codegangsta/cli"
	"github.com/hashicorp/consul/version"
	"github.com/spf13/cobra"
)

var serverWSCmd = &cobra.Command{
	Short: "Start WebSocket Server",
	Long:  `Start WebSocket Server`,
	Use:   "http:ws:server",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		logger.Debug(cmd.Long)

		started := make(chan bool)
		tested := make(chan bool)
		app := cli.NewApp()
		app.Name = "service"
		app.Version = version.Version
		app.Usage = "starts the service."
		app.Action = func(ctx *cli.Context) {
			http.StartServerWS(ctx, started)
		}
		// app.Flags = http.ServerFlags
		go app.Run(os.Args)

		go func() {
			<-started
			logger.Debug("deploy success")
			// tested <- true
		}()
		<-tested
	},
}

func init() {
	RootCmd.AddCommand(serverWSCmd)
}
