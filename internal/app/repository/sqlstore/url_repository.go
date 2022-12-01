package sqlstore

import (
	"context"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type URLRepository struct {
	store *Store
}

func (rep *URLRepository) Add(_ context.Context, url model.ShortURL) error {
	return nil
}

func (rep *URLRepository) GetByCode(_ context.Context, code string) (*model.ShortURL, error) {
	return nil, nil
}

func (rep *URLRepository) DeleteByCode(_ context.Context, code string) error {
	return nil
}
