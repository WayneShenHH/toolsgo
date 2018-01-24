package cmd

import (
	"strconv"

	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/WayneShenHH/toolsgo/services/jusvc"
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

func init() {
	RootCmd.AddCommand(msgCmd)
	RootCmd.AddCommand(juCmd)
}
