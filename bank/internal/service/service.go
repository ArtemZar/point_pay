package service

import (
	"bank/internal/config"
	"bank/internal/model"
	"bank/internal/transport/grpc"
	pb "bank/internal/transport/proto"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	Client *grpc.Client
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		Client: grpc.NewGRPCClient(cfg),
	}
}

func (s Service) CreateAccount(requestBody model.CreateAccountDTO) (*pb.AccountResponse, error) {
	response, err := s.Client.GRPC.CreateAccount(context.Background(), &pb.NewUserRequest{
		Email: requestBody.Email,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s Service) GenerateAddress(requestBody model.UpdateAccountDTO) (*pb.AccountResponse, error) {
	response, err := s.Client.GRPC.GenerateAddress(context.Background(), &pb.NewWalletRequest{
		Id: requestBody.ID,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s Service) Deposit(requestBody model.UpdateAccountDTO) (*pb.AccountResponse, error) {
	response, err := s.Client.GRPC.Deposit(context.Background(), &pb.ChangeBalanceRequest{
		Id:     requestBody.ID,
		Change: requestBody.Amount,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s Service) Withdrawal(requestBody model.UpdateAccountDTO) (*pb.AccountResponse, error) {
	response, err := s.Client.GRPC.Withdrawal(context.Background(), &pb.ChangeBalanceRequest{
		Id:     requestBody.ID,
		Change: requestBody.Amount,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s Service) GetAccounts() (pb.Accounts_GetAccountsClient, error) {
	response, err := s.Client.GRPC.GetAccounts(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return response, nil
}
