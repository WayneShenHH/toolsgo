package cmd

import (
	"strconv"
	"time"

	"github.com/WayneShenHH/toolsgo/tools"
	"github.com/WayneShenHH/toolsgo/tools/timeutil"

	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/WayneShenHH/toolsgo/services/jusvc"
	"github.com/WayneShenHH/toolsgo/services/locksvc"
	"github.com/spf13/cobra"
)

var msgCmd = &cobra.Command{
	Short: "msg cmd",
	Long:  "insert msg cmd",
	Use:   "msg",
	Run: func(cmd *cobra.Command, args []string) {
		repo := repositoryimpl.New()
		ju := jusvc.New(repo)
		switch args[0] {
		case "match":
			ju.InsertMessage("worker:match:message", "msg_match")
		case "offer":
			ju.InsertMessage("worker:offer:message", "msg_offer")
		case "bp":
			ju.InsertMessage("Broadcast:Player", "tmp")
		case "bo":
			ju.InsertMessage("Broadcast:Operator", "tmp")
		case "variant":
			ju.InsertMessage("worker:variant:message", "variant")
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
		repo := repositoryimpl.New()
		ju := jusvc.New(repo)
		ju.CreateJuMatch(uint(i))
	},
}
var txCmd = &cobra.Command{
	Short: "tx cmd",
	Long:  `create tx match`,
	Use:   "tx",
	Run: func(cmd *cobra.Command, args []string) {
		mid := args[0]
		i, _ := strconv.Atoi(mid)
		repo := repositoryimpl.New()
		ju := jusvc.New(repo)
		ju.CreateTxMatch(uint(i))
	},
}
var timerCmd = &cobra.Command{
	Short: "timer cmd",
	Long:  `timer`,
	Use:   "timer",
	Run: func(cmd *cobra.Command, args []string) {
		n := timeutil.TimeToString(time.Now())
		tools.Log(n)
		go locksvc.UsingLockJob("123", "manager")
		go locksvc.UsingLockJob("123", "boss")
		select {}
	},
}

func init() {
	RootCmd.AddCommand(txCmd)
	RootCmd.AddCommand(msgCmd)
	RootCmd.AddCommand(juCmd)
	RootCmd.AddCommand(timerCmd)
}
