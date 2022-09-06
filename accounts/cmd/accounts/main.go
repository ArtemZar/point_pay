package main

import (
	apiGrpc "accounts/internal/api/grpc"
	pb "accounts/internal/api/proto"
	"accounts/internal/config"
	"accounts/internal/db/mongodb"
	"accounts/internal/utils"
	"context"
	"github.com/spf13/viper"
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

	// init configs
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		utils.Logger.Errorf("can't read config from file. error: %v. Will be use default configs", err)
	}
	cfg, err := config.New()
	if err != nil {
		utils.Logger.Fatalf("can't load config with error: %v", err)
	}

	ctx, finish := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	monogodbClient, err := mongodb.NewClient(ctx, cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Database, cfg.DB.AuthDB)
	if err != nil {
		utils.Logger.Fatal(err)
	}

	go func(ctx context.Context) {
		listener, err := net.Listen(cfg.SrvConfig.Network, cfg.SrvConfig.Addr)
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
