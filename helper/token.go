package helper

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateRandomToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
