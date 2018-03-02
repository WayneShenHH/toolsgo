package cmd

import (
	"github.com/WayneShenHH/toolsgo/services/metrixsvc"
	"github.com/spf13/cobra"
)

var metrixCmd = &cobra.Command{
	Short: "metrix server",
	Long:  `metrix watching service healthy`,
	Use:   "metrix",

	Run: func(cmd *cobra.Command, args []string) {
		metrixsvc.MetrixServer()
	},
}
var prometheusCmd = &cobra.Command{
	Short: "prometheusCmd server",
	Long:  `prometheusCmd watching service healthy`,
	Use:   "prometheus",

	Run: func(cmd *cobra.Command, args []string) {
		metrixsvc.PrometheusServer()
		//metrixsvc.Report()
	},
}

func init() {
	RootCmd.AddCommand(prometheusCmd)
}
