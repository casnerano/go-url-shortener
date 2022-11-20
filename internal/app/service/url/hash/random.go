package hash

import (
	"errors"
	"math/rand"
	"time"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var symbols_count = len(symbols)

type Random struct {
	min_len int
	max_len int
}

func (r Random) Generate(string) string {
	l := rand.Intn(r.max_len-r.min_len) + r.min_len
	return r.getRandomString(l)
}

func (r Random) getRandomString(n int) string {
	st := make([]rune, n)
	for i := range st {
		st[i] = symbols[rand.Intn(symbols_count)]
	}
	return string(st)
}

func NewRandom(min, max int) (*Random, error) {
	if min < 1 || max < 1 || min > max {
		return nil, errors.New("invalid arguments")
	}
	return &Random{min_len: min, max_len: max}, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
