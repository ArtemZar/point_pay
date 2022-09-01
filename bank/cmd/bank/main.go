package main

import (
	"bank/internal/api/rest"
	"bank/internal/config"
	"bank/internal/utils"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := utils.InitializeLogger(); err != nil {
		log.Fatalf("Can't initialize logger. Error: %v", err)
	}

	ctx, finish := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func(ctx context.Context) {
		s := rest.NewServer(config.NewCofig())
		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	<-sigCh
	finish()
}