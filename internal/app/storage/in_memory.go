package storage

import (
	"errors"
	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type ShortURLDataStorage = map[string]model.ShortURL

type InMemory struct {
	store ShortURLDataStorage
}

func (m *InMemory) AddURL(url model.ShortURL) error {
	m.store[url.Code] = url
	return nil
}

func (m InMemory) GetURL(code string) (*model.ShortURL, error) {
	url, ok := m.store[code]
	if !ok {
		return nil, errors.New("url not found")
	}
	return &url, nil
}

func NewInMemory() *InMemory {
	return &InMemory{store: make(ShortURLDataStorage)}
}
