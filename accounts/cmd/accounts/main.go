package main

import (
	apiGrpc "accounts/internal/api/grpc"
	pb "accounts/internal/api/proto"
	"accounts/internal/db/mongodb"
	"accounts/internal/utils"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	if err := utils.InitializeLogger(); err != nil {
		log.Fatalf("can't initialize logger. Error: %v", err)
	}

	ctx, finish := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	monogodbClient, err := mongodb.NewClient(ctx, "localhost", "27017", "", "", "account-service", "")
	if err != nil {
		utils.Logger.Fatal(err)
	}

	go func(ctx context.Context) {
		listener, err := net.Listen("tcp", "localhost:8081")
		if err != nil {
			utils.Logger.Fatalf("listener error: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterAccountsServer(s, &apiGrpc.GRPCServer{
			MongoDBClient: monogodbClient,
		})

		if err = s.Serve(listener); err != nil {
			utils.Logger.Info(err)
		}

	}(ctx)

	<-sigCh
	finish()
}
