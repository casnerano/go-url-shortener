package repository

import (
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/storage"
)

type ShortURL struct {
	storage storage.Storage
}

func (s ShortURL) GetURLByCode(code string) (*model.ShortURL, error) {
	return s.storage.GetURL(code)
}

func (s ShortURL) AddURL(url model.ShortURL) error {
	return s.storage.AddURL(url)
}

func NewShortURL(s storage.Storage) *ShortURL {
	return &ShortURL{storage: s}
}
