package model

import "time"

type ShortURL struct {
	ID        int    `json:"id,omitempty"`
	Code      string `json:"code"`
	Original  string `json:"original"`
	UserID    UserID
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func NewShortURL(code, original string) *ShortURL {
	return &ShortURL{
		Code:     code,
		Original: original,
	}
}
