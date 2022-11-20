package repository

import (
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/storage"
)

type ShortUrl struct {
	storage storage.Storage
}

func (s ShortUrl) GetUrlByCode(code string) (*model.ShortUrl, error) {
	return s.storage.GetUrl(code)
}

func (s ShortUrl) AddUrl(url model.ShortUrl) error {
	return s.storage.AddUrl(url)
}

func NewShortUrl(s storage.Storage) *ShortUrl {
	return &ShortUrl{storage: s}
}
