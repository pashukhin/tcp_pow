package handler

import (
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/pashukhin/tcp_pow/pkg/pow"
	"github.com/pashukhin/tcp_pow/pkg/quotes"
	"github.com/pashukhin/tcp_pow/pkg/server"
	"go.uber.org/zap"
)

func NewHandler(logger *zap.Logger, algo string, difficulty int, quotes quotes.Quotes) server.ServerHandler {
	return func(rw io.ReadWriter) {
		// make unique challenge for a client
		challenge := uuid.NewString()

		// setup logger
		logger = logger.With(
			zap.String("algo", algo),
			zap.Int("difficulty", difficulty),
			zap.String("challenge", challenge),
		)

		logger.Info("proof of work task")

		// sending task to client
		fmt.Fprintf(rw, "Proof of Work challenge: %s\n", challenge)
		fmt.Fprintf(rw, "Hash algorithm: %s\n", algo)
		fmt.Fprintf(rw, "Difficulty: %d\n", difficulty)

		// waiting for solution
		solutionBuf := make([]byte, 1024)
		n, err := rw.Read(solutionBuf)
		if err != nil && err != io.EOF {
			logger.Error("Error reading solution", zap.Error(err))
			return
		}
		solution := string(solutionBuf[:n])

		logger = logger.With(zap.String("solution", solution))
		logger.Info("proof of work solution")

		// checking solution
		if pow.VerifyProofOfWork(algo, difficulty, challenge, solution) {
			logger.Info("proof of work verified")
			fmt.Fprintf(rw, quotes.Quote()+"\n")
		} else {
			logger.Info("proof of work failed")
			fmt.Fprintf(rw, "Invalid Proof of Work. Connection closed.\n")
			return
		}
	}
}
