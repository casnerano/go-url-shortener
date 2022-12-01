package memstore

import (
	"context"
	"errors"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type URLRepository struct {
	store *Store
}

func (rep *URLRepository) Add(_ context.Context, url model.ShortURL) error {
	rep.store.shortURLStorage[url.Code] = url
	return nil
}

func (rep *URLRepository) GetByCode(_ context.Context, code string) (*model.ShortURL, error) {
	url, ok := rep.store.shortURLStorage[code]
	if !ok {
		return nil, errors.New("url not found")
	}
	return &url, nil
}

func (rep *URLRepository) DeleteByCode(_ context.Context, code string) error {
	_, ok := rep.store.shortURLStorage[code]
	if !ok {
		return errors.New("url not found")
	}

	delete(rep.store.shortURLStorage, code)
	return nil
}
