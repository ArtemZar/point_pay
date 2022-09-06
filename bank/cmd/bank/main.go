package main

import (
	"bank/internal/api/rest"
	"bank/internal/config"
	"bank/internal/utils"
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := utils.InitializeLogger(); err != nil {
		log.Fatalf("Can't initialize logger. Error: %v", err)
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

	go func(ctx context.Context, cfg *config.Config) {
		s := rest.NewServer(cfg)
		if err := s.Start(); err != nil {
			utils.Logger.Fatal(err)
		}
	}(ctx, cfg)

	<-sigCh
	finish()
}
