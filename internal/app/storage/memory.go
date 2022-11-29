package storage

import (
	"errors"
	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type ShortURLDataStorage = map[string]model.ShortURL

type Memory struct {
	store ShortURLDataStorage
}

func (m *Memory) AddURL(url model.ShortURL) error {
	m.store[url.Code] = url
	return nil
}

func (m *Memory) GetURL(code string) (*model.ShortURL, error) {
	url, ok := m.store[code]
	if !ok {
		return nil, errors.New("url not found")
	}
	return &url, nil
}

func (m *Memory) Reset() {
	m.store = make(ShortURLDataStorage)
}

func NewMemory() *Memory {
	return &Memory{store: make(ShortURLDataStorage)}
}
