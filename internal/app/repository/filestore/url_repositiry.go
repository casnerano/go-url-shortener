package filestore

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

// URLRepository structure for url repository with file store.
type URLRepository struct {
	store *Store
}

// Adding entity.
func (rep *URLRepository) Add(ctx context.Context, url *model.ShortURL) error {
	defer rep.store.Commit(false)

	err := rep.store.memStore.URL().Add(ctx, url)
	if err != nil {
		return err
	}

	_ = rep.store.Write2Buffer(url)
	return nil
}

// Batch adding entities.
func (rep *URLRepository) AddBatch(ctx context.Context, urls []*model.ShortURL) error {
	defer rep.store.Commit(false)

	for _, url := range urls {
		_ = rep.store.memStore.URL().Add(ctx, url)
		_ = rep.store.Write2Buffer(url)
	}

	return nil
}

// Get entity by short code.
func (rep *URLRepository) GetByCode(ctx context.Context, code string) (*model.ShortURL, error) {
	return rep.store.memStore.URL().GetByCode(ctx, code)
}

// Get entity by user uuid and original url.
func (rep *URLRepository) GetByUserUUIDAndOriginal(ctx context.Context, uuid string, original string) (*model.ShortURL, error) {
	return rep.store.memStore.URL().GetByUserUUIDAndOriginal(ctx, uuid, original)
}

// Find entities by user uuid.
func (rep *URLRepository) FindByUserUUID(ctx context.Context, uuid string) ([]*model.ShortURL, error) {
	return rep.store.memStore.URL().FindByUserUUID(ctx, uuid)
}

// Delete entity by short code.
func (rep *URLRepository) DeleteByCode(ctx context.Context, code string, uuid string) error {
	defer rep.store.Commit(true)
	return rep.store.memStore.URL().DeleteByCode(ctx, code, uuid)
}

// Batch delete entities by codes.
func (rep *URLRepository) DeleteBatchByCodes(ctx context.Context, codes []string, uuid string) error {
	defer rep.store.Commit(true)
	return rep.store.memStore.URL().DeleteBatchByCodes(ctx, codes, uuid)
}

// Delete older entities for duration.
func (rep *URLRepository) DeleteOlderRows(ctx context.Context, age time.Duration) error {
	defer rep.store.Commit(true)
	return rep.store.memStore.URL().DeleteOlderRows(ctx, age)
}

// Get total Short URL count
func (rep *URLRepository) GetTotalURLCount(ctx context.Context) (int, error) {
	return rep.store.memStore.URL().GetTotalURLCount(ctx)
}

// Get total User count
func (rep *URLRepository) GetTotalUserCount(ctx context.Context) (int, error) {
	return rep.store.memStore.URL().GetTotalUserCount(ctx)
}
