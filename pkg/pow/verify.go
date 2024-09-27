package pow

import "strings"

func VerifyProofOfWork(algo string, difficulty int, challenge string, solution string) bool {
	hash := CalculateHash(algo, challenge+strings.TrimSpace(solution))
	return strings.HasPrefix(hash, strings.Repeat("0", difficulty))
}
