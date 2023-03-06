package hasher

import (
	"crypto/sha256"
)

type Unique struct{}

func NewUnique() Hash {
	return &Unique{}
}

func (u Unique) Generate(url string) string {
	sha := sha256.Sum256([]byte(url))
	return string(sha[:])
}
