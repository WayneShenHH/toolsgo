package integration

import (
	"context"
	"fmt"

	"git.b5k.io/fsbs/libgo/models/entities"
	"git.b5k.io/fsbs/libgo/store"

	"git.b5k.io/fsbs/libgo/pb"
)

// IdentityServer
func (s *suiteContext) SSO(context.Context, *pb.SSORequest) (*pb.SSOResponse, error) {
	return nil, nil
}

func (s *suiteContext) Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *suiteContext) Logout(context.Context, *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, nil
}

func (s *suiteContext) GetAgent(context.Context, *pb.GetAgentRequest) (*pb.Agent, error) {
	return nil, nil
}

func (s *suiteContext) GetPlayerDetail(context.Context, *pb.GetPlayerDetailRequest) (*pb.GetPlayerDetailResponse, error) {
	return nil, nil
}

func (s *suiteContext) GetPlayerList(ctx context.Context, req *pb.GetPlayerListRequest) (*pb.GetPlayerListResponse, error) {
	return nil, nil
}

func (s *suiteContext) GetAccountDetail(context.Context, *pb.GetAccountDetailRequest) (*pb.AccountDetail, error) {
	return nil, nil
}

// GameServer
func (s *suiteContext) GetMatches(context.Context, *pb.GetMatchesRequest) (*pb.GetMatchesResponse, error) {
	return nil, nil
}

// TransactionServer
func (s *suiteContext) Purchase(context.Context, *pb.Order) (*pb.PurchaseResponse, error) {
	return nil, nil
}

func (s *suiteContext) Deposit(context.Context, *pb.XferInOutRequest) (*pb.Wallet, error) {
	return nil, nil
}

func (s *suiteContext) Withdraw(context.Context, *pb.XferInOutRequest) (*pb.Wallet, error) {
	return nil, nil
}

func (s *suiteContext) SettleOrder(context.Context, *pb.SettleOrderRequest) (*pb.SettleOrderResponse, error) {
	return nil, nil
}

func (s *suiteContext) SettleOrderItem(context.Context, *pb.SettleOrderItemRequest) (*pb.SettleOrderItemResponse, error) {
	return nil, nil
}

// XferHistoryServer
func (s *suiteContext) GetHistories(context.Context, *pb.GetHistoriesRequest) (*pb.Histories, error) {
	return nil, nil
}

func (s *suiteContext) GetSimpleHistories(context.Context, *pb.GetSimpleHistoriesRequest) (*pb.SimpleHistories, error) {
	return nil, nil
}

func (s *suiteContext) startMlmServer4Order() error {
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterGameServer(grpcServer, s)
		pb.RegisterIdentityServer(grpcServer, s)
		pb.RegisterTransactionServer(grpcServer, s)
		pb.RegisterXferHistoryServer(grpcServer, s)

		l, err := net.Listen("tcp", environment.Setting.GRPC.MLMAPI)
		if err != nil {
			panic(err)
		}

		fmt.Println("Serving grpc on", environment.Setting.GRPC.MLMAPI)
		grpcServer.Serve(l)
	}()
	return nil
}