package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pashukhin/tcp_pow/internal/client/request"
	"github.com/pashukhin/tcp_pow/pkg/client"
	"go.uber.org/zap"
)

func main() {
	serverAddr := flag.String("addr", "localhost:8080", "server address")
	numConnections := flag.Int("nconns", 1, "number of parallel connections")
	numRequests := flag.Int("nreqs", 1, "number of requests per connection")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	logger = logger.With(
		zap.String("addr", *serverAddr),
		zap.Int("nconns", *numConnections),
		zap.Int("nreqs", *numRequests),
	)
	logger.Info("Starting client")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < *numConnections; i++ {
		client := client.New(logger, *serverAddr, request.Request)
		wg.Add(1)
		go client.Run(ctx, &wg, *numRequests)
	}

	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		done <- true
	}()

L:
	for {
		select {
		case <-stop:
			logger.Info("Shutting down...")
			cancel()
		case <-done:
			logger.Info("All clients finished.")
			break L
		}
	}
}
