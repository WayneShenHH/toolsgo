package cmd

import (
	"strconv"

	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/WayneShenHH/toolsgo/services/jusvc"
	"github.com/WayneShenHH/toolsgo/services/txsvc"
	"github.com/spf13/cobra"
)

var msgCmd = &cobra.Command{
	Short: "msg cmd",
	Long:  "insert msg cmd",
	Use:   "msg",
	Run: func(cmd *cobra.Command, args []string) {
		repo := repositoryimpl.New(false)
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
		repo := repositoryimpl.New(false)
		ju := jusvc.New(repo)
		if clearFlag {
			ju.Clear()
		}
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
		repo := repositoryimpl.New(false)
		ju := jusvc.New(repo)
		if clearFlag {
			ju.Clear()
		}
		ju.CreateTxMatch(uint(i))
	},
}
var txAdaptorCmd = &cobra.Command{
	Short: "cmd",
	Long:  `tx message`,
	Use:   "txmsg",
	Run: func(cmd *cobra.Command, args []string) {
		repo := repositoryimpl.New(false)
		tx := txsvc.New(repo)
		tx.GetTxMsg(matchID)
	},
}
var (
	clearFlag bool
	matchID   uint
)

func init() {
	txCmd.Flags().BoolVarP(&clearFlag, "clear", "c", false, "clear data")
	juCmd.Flags().BoolVarP(&clearFlag, "clear", "c", false, "clear data")
	RootCmd.AddCommand(txCmd)
	RootCmd.AddCommand(msgCmd)
	RootCmd.AddCommand(juCmd)
	txAdaptorCmd.Flags().UintVarP(&matchID, "matchid", "m", 0, "find tx origin message")
	RootCmd.AddCommand(txAdaptorCmd)
}
