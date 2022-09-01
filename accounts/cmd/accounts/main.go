package main

import (
	pb "accounts/internal/api/gen/proto"
	apiGrpc "accounts/internal/api/grpc"
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

	//monogodbClient, err := mongodb.NewClient(ctx, "localhost", "27017", "", "", "account-service", "")
	//if err != nil {
	//	utils.Logger.Fatal(err)
	//}
	//newAccount := model.Account{
	//	ID:      "",
	//	Email:   "knopa.tan@gmail.com",
	//	Balance: "0",
	//}

	//updateBalance := model.Account{
	//	Email:   "zarubinaav178@gmail.com",
	//	Balance: "5000",
	//}
	//storage := mongodb.NewStorage(monogodbClient, "accounts", utils.Logger)
	//storage.Create(ctx, newAccount)
	//_, err = storage.Update(ctx, updateBalance)
	//if err != nil {
	//	utils.Logger.Info("update error ", err)
	//}

	go func(ctx context.Context) {
		listener, err := net.Listen("tcp", "localhost:8081")
		if err != nil {
			utils.Logger.Fatalf("listener error: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterAccountsServer(s, &apiGrpc.GRPCServer{})

		if err = s.Serve(listener); err != nil {
			utils.Logger.Info(err)
		}

		//s := rest.NewServer(config.NewCofig())
		//if err := s.Start(); err != nil {
		//	log.Fatal(err)
		//}
	}(ctx)

	<-sigCh
	finish()
}
