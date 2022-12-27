package repository

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type URLRepository interface {
	Add(ctx context.Context, url *model.ShortURL) error
	GetByCode(ctx context.Context, code string) (*model.ShortURL, error)
	FindByUser(ctx context.Context, uid model.UserID) ([]*model.ShortURL, error)
	DeleteByCode(ctx context.Context, code string) error
	DeleteOlderRows(ctx context.Context, age time.Duration) error
}
