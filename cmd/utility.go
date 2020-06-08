package cmd

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/tools/timeutil"

	"github.com/spf13/cobra"
)

var ts int64
var timestampCmd = &cobra.Command{
	Short: "utility timestamp",
	Long:  "timestamp transfer",
	Use:   "ts",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
		fmt.Println(cmd.Long)

		time := timeutil.StampToTime(ts)

		if time.Year() < 2000 {
			time = timeutil.StampToTime(ts * 1000)
		}
		fmt.Println(time)
	},
}

func init() {
	timestampCmd.Flags().Int64VarP(&ts, "ts", "t", 0, "timestamp")
	RootCmd.AddCommand(timestampCmd)
}
