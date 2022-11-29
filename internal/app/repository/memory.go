package repository

import (
	"context"
	"errors"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type ShortURLDataStorage = map[string]model.ShortURL

type Memory struct {
	store ShortURLDataStorage
}

func (m *Memory) Add(_ context.Context, url model.ShortURL) error {
	m.store[url.Code] = url
	return nil
}

func (m *Memory) GetByCode(_ context.Context, code string) (*model.ShortURL, error) {
	url, ok := m.store[code]
	if !ok {
		return nil, errors.New("url not found")
	}
	return &url, nil
}

func (m *Memory) DeleteByCode(_ context.Context, code string) error {
	_, ok := m.store[code]
	if !ok {
		return errors.New("url not found")
	}

	delete(m.store, code)
	return nil
}

func NewMemory() *Memory {
	return &Memory{store: make(ShortURLDataStorage)}
}
