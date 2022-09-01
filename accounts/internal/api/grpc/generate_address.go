package grpc

import (
	pb "accounts/internal/api/gen/proto"
	"accounts/internal/db/model"
	"accounts/internal/db/mongodb"
	"accounts/internal/utils"
	"context"
	"encoding/binary"
	"github.com/google/uuid"
)

func (s *GRPCServer) GenerateAddress(ctx context.Context, in *pb.NewWalletRequest) (*pb.AccountResponse, error) {
	uuid := uuid.New()
	walletAddress := binary.BigEndian.Uint64(uuid[:16])

	monogodbClient, err := mongodb.NewClient(ctx, "localhost", "27017", "", "", "account-service", "")
	if err != nil {
		utils.Logger.Fatal(err)
	}

	updateWallet := model.Account{
		ID:       in.Id,
		WalletID: walletAddress,
	}

	storage := mongodb.NewStorage(monogodbClient, "accounts", utils.Logger)
	err = storage.Update(ctx, updateWallet)
	if err != nil {
		//TODO return error
		utils.Logger.Info("update error ", err)
	}

	return &pb.AccountResponse{Id: in.Id, WalletId: walletAddress}, nil
}
