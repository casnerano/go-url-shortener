package repository

import (
	"context"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type URLRepository interface {
	Add(ctx context.Context, url model.ShortURL) error
	GetByCode(ctx context.Context, code string) (*model.ShortURL, error)
	DeleteByCode(ctx context.Context, code string) error
}
