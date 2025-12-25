package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSHA256(input []byte) string {
	h := sha256.New()

	// Write the input data (as a byte slice) to the hash object
	h.Write(input)

	// Get the finalized hash result as a byte slice
	hashBytes := h.Sum(nil)

	// Encode the byte slice to a human-readable hexadecimal string
	return hex.EncodeToString(hashBytes)
}