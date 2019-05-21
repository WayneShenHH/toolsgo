package rpc

import (
	"github.com/WayneShenHH/toolsgo/app"
	"github.com/WayneShenHH/toolsgo/module/logger"
	"github.com/WayneShenHH/toolsgo/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Client gRPC客端(目前用於連接MLM)
type Client interface {
	LoginMLM(username, password, IP string) (token string, _ error)
}

type client struct{}

var c = &client{}

// GetClient 取得gRPC客端服務
func GetClient() Client { return c }

// 若 err 不為 nil 時透過 Logger.Warning 將其印出
func alert(err error) {
	if err != nil {
		logger.Warning(err)
	}
}

// LoginMLM auth mlm, getting token if success
func (c *client) LoginMLM(username, password, IP string) (string, error) {
	conn, err := connect()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	svc := pb.NewIdentityClient(conn)
	req := &pb.LoginRequest{
		Username: username,
		Password: password,
		IP:       IP,
	}
	res, err := svc.Login(context.Background(), req)
	if err != nil {
		return "", err
	}

	return res.Token, err
}

func connect() (*grpc.ClientConn, error) {
	return grpc.Dial(app.Setting.GRPC.MLMAPI, grpc.WithInsecure())
}
func convertUintPtr(u *uint) int64 {
	i := int64(0)
	if u != nil {
		i = int64(*u)
	}
	return i
}
