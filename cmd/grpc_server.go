package cmd

import (
	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/WayneShenHH/toolsgo/transport/rpc"
	"github.com/spf13/cobra"
)

var grpcServerCmd = &cobra.Command{
	Use:   "grpc:server",
	Short: "Start gRPC Server",
	Run: func(_ *cobra.Command, _ []string) {
		logger.Info("[gRPC Server] Server start")
		rpc.Run()
	},
}

func init() {
	RootCmd.AddCommand(grpcServerCmd)
}
