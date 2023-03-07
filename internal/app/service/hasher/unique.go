package hasher

import (
	"crypto/sha256"
	"encoding/hex"
)

// Unique - structure for unique hasher
type Unique struct{}

// NewUnique - constructor.
func NewUnique() Hash {
	return &Unique{}
}

// Generate return unique string
func (u Unique) Generate(url string) string {
	sha := sha256.Sum256([]byte(url))
	return hex.EncodeToString(sha[:])
}
