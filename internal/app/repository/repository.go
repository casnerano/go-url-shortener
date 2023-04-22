// Package with repositories for working with various repositories.
package repository

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

// URLRepository describes how to work with `ShortURL`.
type URLRepository interface {
	// Adding entity.
	Add(ctx context.Context, url *model.ShortURL) error
	// Batch adding entities.
	AddBatch(ctx context.Context, urls []*model.ShortURL) error
	// Get entity by short code.
	GetByCode(ctx context.Context, code string) (*model.ShortURL, error)
	// Get entity by user uuid and original url.
	GetByUserUUIDAndOriginal(ctx context.Context, uuid string, original string) (*model.ShortURL, error)
	// Find entities by user uuid.
	FindByUserUUID(ctx context.Context, uuid string) ([]*model.ShortURL, error)
	// Delete entity by short code.
	DeleteByCode(ctx context.Context, code string, uuid string) error
	// Batch delete entities by codes.
	DeleteBatchByCodes(ctx context.Context, codes []string, uuid string) error
	// Delete older entities for duration.
	DeleteOlderRows(ctx context.Context, age time.Duration) error

	// Get total Short URL count
	GetTotalURLCount(ctx context.Context) (int, error)

	// Get total User count
	GetTotalUserCount(ctx context.Context) (int, error)
}
