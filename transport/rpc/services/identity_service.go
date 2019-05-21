package services

/*
	若要啟動服務必須把 pb 內用到的方法實作完成
*/

import (
	"context"

	"github.com/WayneShenHH/toolsgo/pb"
	"github.com/WayneShenHH/toolsgo/repository"
)

type service struct{ repository.Context }

// NewIdentityServer 初始化 OrderInfoServer
func NewIdentityServer(ctx repository.Context) pb.IdentityServer {
	return &service{Context: ctx}
}

// Login
func (*service) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}
