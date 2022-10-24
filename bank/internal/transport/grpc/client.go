package grpc

import (
	"bank/internal/config"
	pb "bank/internal/transport/proto"
	"bank/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	GRPC pb.AccountsClient
}

func NewGRPCClient(cfg *config.Config) *Client {
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial(cfg.GrpcClientConf.Target, grpc.WithTransportCredentials(creds))
	if err != nil {
		utils.Logger.Fatalf("GRPC client error: %v", err)
	}

	return &Client{GRPC: pb.NewAccountsClient(conn)}
}
