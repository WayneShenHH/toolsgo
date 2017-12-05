package cmd

import (
	"fmt"
	"os"

	"github.com/WayneShenHH/toolsgo/transport/http"
	"github.com/codegangsta/cli"
	"github.com/hashicorp/consul/version"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Use: "server",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
		fmt.Println(cmd.Long)

		started := make(chan bool)
		tested := make(chan bool)
		app := cli.NewApp()
		app.Name = "service"
		app.Version = version.Version
		app.Usage = "starts the service."
		app.Action = func(ctx *cli.Context) {
			http.StartServer(ctx, started)
		}

		app.Flags = http.ServerFlags
		go app.Run(os.Args)

		go func() {
			<-started
			fmt.Println("deploy success")
		}()
		<-tested
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}
