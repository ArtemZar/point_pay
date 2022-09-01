package grpc

import (
	pb "accounts/internal/api/gen/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) GetAccounts(*emptypb.Empty, pb.Accounts_GetAccountsServer) error {
	return nil
}
