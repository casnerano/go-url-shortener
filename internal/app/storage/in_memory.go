package storage

import (
	"errors"
	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type ShortUrlDataStorage = map[string]model.ShortUrl

type InMemory struct {
	store ShortUrlDataStorage
}

func (m *InMemory) AddUrl(url model.ShortUrl) error {
	m.store[url.Code] = url
	return nil
}

func (m InMemory) GetUrl(code string) (*model.ShortUrl, error) {
	url, ok := m.store[code]
	if !ok {
		return nil, errors.New("url not found")
	}
	return &url, nil
}

func NewInMemory() *InMemory {
	return &InMemory{store: make(ShortUrlDataStorage)}
}
