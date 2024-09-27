package pow

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

func SolveProofOfWork(ctx context.Context, algo string, difficulty int, challenge string) (string, error) {
	var solution string
	prefix := strings.Repeat("0", difficulty)
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return "", errors.New("interrupted")
		default:
			solution = strconv.Itoa(i)
			hash := CalculateHash(algo, challenge+solution)
			if strings.HasPrefix(hash, prefix) {
				return solution, nil
			}
		}
	}
}
