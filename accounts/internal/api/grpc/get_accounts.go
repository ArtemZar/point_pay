package grpc

import (
	pb "accounts/internal/api/proto"
	"accounts/internal/db/mongodb"
	"accounts/internal/utils"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) GetAccounts(_ *emptypb.Empty, stream pb.Accounts_GetAccountsServer) error {

	storage := mongodb.NewStorage(s.MongoDBClient, "accounts", utils.Logger)

	allAccs, _ := storage.Find(context.Background())

	//nolint:errcheck // ok
	for _, val := range allAccs {
		stream.Send(&pb.AccountResponse{
			Id:       val.ID,
			WalletId: val.WalletID,
			Balance:  val.Balance,
		})

	}

	return nil
}
