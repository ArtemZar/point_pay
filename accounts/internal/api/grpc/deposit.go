package grpc

import (
	pb "accounts/internal/api/gen/proto"
	"accounts/internal/db/model"
	"accounts/internal/db/mongodb"
	"accounts/internal/service"
	"accounts/internal/utils"
	"context"
)

func (s *GRPCServer) Deposit(ctx context.Context, in *pb.ChangeBalanceRequest) (*pb.AccountResponse, error) {
	monogodbClient, err := mongodb.NewClient(ctx, "localhost", "27017", "", "", "account-service", "")
	if err != nil {
		utils.Logger.Fatal(err)
	}

	storage := mongodb.NewStorage(monogodbClient, "accounts", utils.Logger)

	sourceAcc, _ := storage.GetOne(ctx, model.Account{ID: in.Id})

	newBalance, _ := service.ChangeBalance(sourceAcc.Balance, in.Change, service.Deposite{})

	updateBalance := model.Account{
		ID:      in.Id,
		Balance: newBalance,
	}

	err = storage.Update(ctx, updateBalance)
	if err != nil {
		//TODO return error
		utils.Logger.Info("update error ", err)
	}

	return &pb.AccountResponse{Id: in.Id, Balance: newBalance}, nil
}