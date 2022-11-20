package model

import "time"

type ShortUrl struct {
	Code      string        `json:"code"`
	Original  string        `json:"original"`
	CreatedAt time.Time     `json:"createdAt"`
	LifeTime  time.Duration `json:"lifeTime"`
}

func NewShortUrl(code, original string, lifeTime time.Duration) *ShortUrl {
	return &ShortUrl{
		Code:      code,
		Original:  original,
		CreatedAt: time.Now(),
		LifeTime:  lifeTime,
	}
}
