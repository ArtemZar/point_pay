package grpc

import (
	pb "accounts/internal/api/proto"
)

type handlers struct {
	pb.AccountsServer
	//service service
}

func NewHandlers() *handlers {
	return &handlers{}
}
