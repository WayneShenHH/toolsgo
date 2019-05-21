package rpc

import (
	"fmt"
	"net"

	"github.com/WayneShenHH/toolsgo/transport/rpc/services"

	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/pb"
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"

	"google.golang.org/grpc"
)

// Run 啟動 gRPC Server
func Run() {
	listener, err := net.Listen("tcp", app.Setting.GRPC.Server)
	if err != nil {
		panic(err)
	}
	ctx := repository.Context{Repository: repositoryimpl.New()}
	g := grpc.NewServer()

	svc := services.NewIdentityServer(ctx)
	pb.RegisterIdentityServer(g, svc)

	fmt.Printf("[gRPC Server] server stop: %v\n", g.Serve(listener))
}
