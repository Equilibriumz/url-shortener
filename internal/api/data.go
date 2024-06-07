package api

import (
	"crypto/rand"
	"math/big"
)

const (
	keyLength = 8

	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateKey() string {
	result := make([]byte, keyLength)

	for i := 0; i < keyLength; i++ {
		charIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[charIndex.Int64()]
	}

	return string(result)
}
