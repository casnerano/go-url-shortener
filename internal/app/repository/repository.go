package repository

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type URLRepository interface {
	Add(ctx context.Context, url *model.ShortURL) error
	AddBatch(ctx context.Context, urls []*model.ShortURL) error
	GetByCode(ctx context.Context, code string) (*model.ShortURL, error)
	GetByUserUUIDAndOriginal(ctx context.Context, uuid string, original string) (*model.ShortURL, error)
	FindByUserUUID(ctx context.Context, uuid string) ([]*model.ShortURL, error)
	DeleteByCode(ctx context.Context, code string, uuid string) error
	DeleteBatchByCodes(ctx context.Context, codes []string, uuid string) error
	DeleteOlderRows(ctx context.Context, age time.Duration) error
}
