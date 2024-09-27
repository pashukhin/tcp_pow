package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pashukhin/tcp_pow/internal/server/handler"
	"github.com/pashukhin/tcp_pow/pkg/quotes"
	"github.com/pashukhin/tcp_pow/pkg/server"

	"go.uber.org/zap"
)

func main() {
	algo := flag.String("algo", "SHA256", "hash algorithm")
	difficulty := flag.Int("diff", 4, "difficulty")
	quotesPath := flag.String("quotes", "", "difficulty")
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	if *quotesPath == "" {
		logger.Error("quoes flag is mandatory")
		os.Exit(1)
	}

	logger.Info("Reading quotes...")
	quotes, err := quotes.New(*quotesPath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Starting server...")
	s, err := server.New(
		logger,
		":8080",
		handler.NewHandler(logger, *algo, *difficulty, quotes),
	)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	s.Start()
	logger.Info("Server is ready to work")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	logger.Info("Shutting down server...")
	s.Stop()
	logger.Info("Server stopped.")
}
