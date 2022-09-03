package grpc

import (
	pb "accounts/internal/api/proto"
	"accounts/internal/db/model"
	"accounts/internal/db/mongodb"
	"accounts/internal/utils"
	"context"
)

func (s *GRPCServer) CreateAccount(ctx context.Context, in *pb.NewUserRequest) (*pb.AccountResponse, error) {

	newAccount := model.Account{
		ID:      "",
		Email:   in.Email,
		Balance: "0",
	}

	storage := mongodb.NewStorage(s.MongoDBClient, "accounts", utils.Logger)
	accountID, _ := storage.Create(ctx, newAccount)

	return &pb.AccountResponse{Id: accountID}, nil
}
