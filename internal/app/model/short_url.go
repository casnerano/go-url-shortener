package model

import "time"

type ShortURL struct {
	Code      string    `json:"code"`
	Original  string    `json:"original"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewShortURL(code, original string) *ShortURL {
	return &ShortURL{
		Code:     code,
		Original: original,
	}
}
