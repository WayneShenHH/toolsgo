package cmd

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var showConfig = &cobra.Command{
	Use:   "config",
	Short: "show setting",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		d, _ := yaml.Marshal(&app.Setting)

		fmt.Printf("--- m dump:\n%s\n\n", string(d))
	}, 
}

func init() {
	RootCmd.AddCommand(showConfig)
}
