package request

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/pashukhin/tcp_pow/pkg/pow"
	"go.uber.org/zap"
)

func Request(ctx context.Context, logger *zap.Logger, rw io.ReadWriter) {
	reader := bufio.NewReader(rw)

	challengeLine, err := reader.ReadString('\n')
	if err != nil {
		logger.Error("Error reading challenge", zap.Error(err))
		return
	}
	algoLine, err := reader.ReadString('\n')
	if err != nil {
		logger.Error("Error reading algo", zap.Error(err))
		return
	}
	difficultyLine, err := reader.ReadString('\n')
	if err != nil {
		logger.Error("Error reading difficulty", zap.Error(err))
		return
	}

	challenge := strings.TrimSpace(strings.Split(challengeLine, ": ")[1])
	algo := strings.TrimSpace(strings.Split(algoLine, ": ")[1])
	difficulty, err := strconv.Atoi(strings.TrimSpace(strings.Split(difficultyLine, ": ")[1]))
	if err != nil {
		logger.Error("Error converting difficulty", zap.Error(err))
		return
	}

	start := time.Now()
	solution, err := pow.SolveProofOfWork(ctx, algo, difficulty, challenge)
	if err != nil {
		logger.Error("Error solwing POW", zap.Error(err))
	}
	logger.Info("time spent", zap.Duration("spent", time.Since(start)))

	fmt.Fprintf(rw, solution+"\n")

	response, err := reader.ReadString('\n')
	if err != nil {
		logger.Error("Error reading difficulty", zap.Error(err))
		return
	}
	logger.Info("Server response", zap.String("response", response))
}
