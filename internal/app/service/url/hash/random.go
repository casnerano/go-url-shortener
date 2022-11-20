package hash

import (
	"errors"
	"math/rand"
	"time"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var symbolsCount = len(symbols)

type Random struct {
	minLen int
	maxLen int
}

func (r Random) Generate(string) string {
	l := rand.Intn(r.maxLen-r.minLen) + r.minLen
	return r.getRandomString(l)
}

func (r Random) getRandomString(n int) string {
	st := make([]rune, n)
	for i := range st {
		st[i] = symbols[rand.Intn(symbolsCount)]
	}
	return string(st)
}

func NewRandom(min, max int) (*Random, error) {
	if min < 1 || max < 1 || min > max {
		return nil, errors.New("invalid arguments")
	}
	return &Random{minLen: min, maxLen: max}, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
