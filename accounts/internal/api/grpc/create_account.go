package grpc

import (
	pb "accounts/internal/api/gen/proto"
	"accounts/internal/db/model"
	"accounts/internal/db/mongodb"
	"accounts/internal/utils"
	"context"
)

func (s *GRPCServer) CreateAccount(ctx context.Context, in *pb.NewUserRequest) (*pb.AccountResponse, error) {
	monogodbClient, err := mongodb.NewClient(ctx, "localhost", "27017", "", "", "account-service", "")
	if err != nil {
		utils.Logger.Fatal(err)
	}

	newAccount := model.Account{
		ID:      "",
		Email:   in.Email,
		Balance: "0",
	}

	storage := mongodb.NewStorage(monogodbClient, "accounts", utils.Logger)
	accountID, _ := storage.Create(ctx, newAccount)

	return &pb.AccountResponse{Id: accountID}, nil
}
