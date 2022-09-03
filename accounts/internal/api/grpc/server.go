package grpc

import (
	pb "accounts/internal/api/proto"
	"go.mongodb.org/mongo-driver/mongo"
)

type GRPCServer struct {
	pb.UnimplementedAccountsServer
	MongoDBClient *mongo.Database
}
