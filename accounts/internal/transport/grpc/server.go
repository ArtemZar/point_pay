package transport

import (
	"accounts/internal/model"
	pb "accounts/internal/transport/proto"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type useCaseService interface {
	CreateAccount(ctx context.Context, email string) (accID string, err error)
	GenerateAddress(ctx context.Context, accID string) (walletID uint64, err error)
	Deposit(ctx context.Context, accID, amountOfChange string) (newBalance string, err error)
	Withdrawal(ctx context.Context, accID, amountOfChange string) (newBalance string, err error)
	GetAccounts() (acc []model.Account, err error)
}

type AccountsSrv struct {
	pb.UnimplementedAccountsServer
	UseCase useCaseService
}

func New(uc useCaseService) *AccountsSrv {
	return &AccountsSrv{UseCase: uc}
}

func (as *AccountsSrv) CreateAccount(ctx context.Context, in *pb.NewUserRequest) (*pb.AccountResponse, error) {
	accountID, err := as.UseCase.CreateAccount(ctx, in.Email)
	if err != nil {
		return &pb.AccountResponse{Id: ""}, fmt.Errorf("faild create account with error %v", err)
	}
	return &pb.AccountResponse{Id: accountID}, nil
}

func (as *AccountsSrv) Deposit(ctx context.Context, in *pb.ChangeBalanceRequest) (*pb.AccountResponse, error) {

	newBalance, err := as.UseCase.Deposit(ctx, in.Id, in.Change)
	if err != nil {
		return &pb.AccountResponse{Id: in.Id}, fmt.Errorf("deposit faild with error %v", err)
	}

	return &pb.AccountResponse{Id: in.Id, Balance: newBalance}, nil
}

func (as *AccountsSrv) GenerateAddress(ctx context.Context, in *pb.NewWalletRequest) (*pb.AccountResponse, error) {
	walletAddress, err := as.UseCase.GenerateAddress(ctx, in.Id)
	if err != nil {
		return &pb.AccountResponse{Id: in.Id}, err
	}

	return &pb.AccountResponse{Id: in.Id, WalletId: walletAddress}, nil
}

func (as *AccountsSrv) Withdrawal(ctx context.Context, in *pb.ChangeBalanceRequest) (*pb.AccountResponse, error) {
	newBalance, err := as.UseCase.Withdrawal(ctx, in.Id, in.Change)
	if err != nil {
		return &pb.AccountResponse{Id: in.Id}, fmt.Errorf("withdrawal faild with error %v", err)
	}

	return &pb.AccountResponse{Id: in.Id, Balance: newBalance}, nil
}

func (as *AccountsSrv) GetAccounts(_ *emptypb.Empty, stream pb.Accounts_GetAccountsServer) error {

	allAccs, err := as.UseCase.GetAccounts()
	if err != nil {
		return fmt.Errorf("faild get accounts with error %v", err)
	}
	for _, val := range allAccs {
		stream.Send(&pb.AccountResponse{
			Id:       val.ID,
			WalletId: uint64(val.WalletID),
			Balance:  val.Balance,
		})

	}

	return nil
}
