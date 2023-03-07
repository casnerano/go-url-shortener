package hasher

import (
	"errors"
	"math/rand"
	"time"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var symbolsCount = len(symbols)

// Struct for random hasher
type Random struct {
	minLen int
	maxLen int
}

// NewRandom returns an object for generating random links, of arbitrary length.
// The `min` and `max` parameters define the short link length limits.
func NewRandom(min, max int) (Hash, error) {
	if min < 1 || max < 1 || min > max {
		return nil, errors.New("invalid arguments")
	}
	return &Random{minLen: min, maxLen: max}, nil
}

// Generate return random string
func (r Random) Generate(_ string) string {
	l := rand.Intn(r.maxLen-r.minLen+1) + r.minLen
	return r.getRandomString(l)
}

func (r Random) getRandomString(n int) string {
	st := make([]rune, n)
	for i := range st {
		st[i] = symbols[rand.Intn(symbolsCount)]
	}
	return string(st)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
