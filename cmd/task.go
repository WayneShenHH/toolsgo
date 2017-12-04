package cmd

import (
	"strconv"

	"github.com/WayneShenHH/toolsgo/services"
	"github.com/spf13/cobra"
)

var msgCmd = &cobra.Command{
	Short: "msg cmd",
	Long:  "insert msg cmd",
	Use:   "msg",
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "match":
			services.InsertMessage("worker:match:message", "match_point")
		case "offer":
			services.InsertMessage("worker:offer:message", "offer_point")
		}
	},
}

var juCmd = &cobra.Command{
	Short: "ju cmd",
	Long:  `create ju match`,
	Use:   "ju",
	Run: func(cmd *cobra.Command, args []string) {
		mid := args[0]
		i, _ := strconv.Atoi(mid)
		services.CreateJuMatch(uint(i))
	},
}

func init() {
	RootCmd.AddCommand(msgCmd)
	RootCmd.AddCommand(juCmd)
}
