package grpc

import (
	pb "bank/internal/api/gen/proto"
	"bank/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Client pb.AccountsClient
}

func NewGRPCClient() *GRPCClient {
	creds := insecure.NewCredentials()
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		utils.Logger.Infof("GRPC client error: %v", err)
	}

	return &GRPCClient{Client: pb.NewAccountsClient(conn)}
	//client := pb.NewAccountsClient(conn)

	//resp, err := client.Test(context.Background(), &pb.Request{
	//	X: 3,
	//	Y: 5,
	//})
}
