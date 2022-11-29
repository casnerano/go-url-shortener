package model

import "time"

type ShortURL struct {
	Code      string        `json:"code"`
	Original  string        `json:"original"`
	CreatedAt time.Time     `json:"createdAt"`
	LifeTime  time.Duration `json:"lifeTime"`
}

func NewShortURL(code, original string, lifeTime time.Duration) *ShortURL {
	return &ShortURL{
		Code:      code,
		Original:  original,
		CreatedAt: time.Now(),
		LifeTime:  lifeTime,
	}
}
