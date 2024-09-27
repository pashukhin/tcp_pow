package pow

import (
	"crypto/sha256"
	"crypto/sha512"
	"crypto/md5"
	"encoding/hex"
)

func CalculateHash(algo string, data string) string {
	switch algo {
	case "MD5":
		hash := md5.Sum([]byte(data))
		return hex.EncodeToString(hash[:])
	case "SHA512":
		hash := sha512.Sum512([]byte(data))
		return hex.EncodeToString(hash[:])
	default:
		hash := sha256.Sum256([]byte(data))
		return hex.EncodeToString(hash[:])
	}
}
