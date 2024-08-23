package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Generates a random string of a specified length.
func GenerateRandomToken(length int) (string, error) {
	
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var token strings.Builder
	
	token.Grow(length)

	for i := 0; i < length; i++ {

		// Generate a random index to select a character from charset
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		token.WriteByte(charset[randomIndex.Int64()])
	}

	return token.String(), nil
}
