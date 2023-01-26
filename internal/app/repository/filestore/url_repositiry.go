package filestore

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type URLRepository struct {
	store *Store
}

func (rep *URLRepository) Add(ctx context.Context, url *model.ShortURL) error {
	defer rep.store.Commit(false)

	err := rep.store.memStore.URL().Add(ctx, url)
	if err != nil {
		return err
	}

	_ = rep.store.Write2Buffer(url)
	return nil
}

func (rep *URLRepository) AddBatch(ctx context.Context, urls []*model.ShortURL) error {
	defer rep.store.Commit(false)

	for _, url := range urls {
		_ = rep.store.memStore.URL().Add(ctx, url)
		_ = rep.store.Write2Buffer(url)
	}

	return nil
}

func (rep *URLRepository) GetByCode(ctx context.Context, code string) (*model.ShortURL, error) {
	return rep.store.memStore.URL().GetByCode(ctx, code)
}

func (rep *URLRepository) GetByUserUUIDAndOriginal(ctx context.Context, uuid string, original string) (*model.ShortURL, error) {
	return rep.store.memStore.URL().GetByUserUUIDAndOriginal(ctx, uuid, original)
}

func (rep *URLRepository) FindByUserUUID(ctx context.Context, uuid string) ([]*model.ShortURL, error) {
	return rep.store.memStore.URL().FindByUserUUID(ctx, uuid)
}

func (rep *URLRepository) DeleteByCode(ctx context.Context, code string, uuid string) error {
	defer rep.store.Commit(true)
	return rep.store.memStore.URL().DeleteByCode(ctx, code, uuid)
}

func (rep *URLRepository) DeleteBatchByCodes(ctx context.Context, codes []string, uuid string) error {
	defer rep.store.Commit(true)
	return rep.store.memStore.URL().DeleteBatchByCodes(ctx, codes, uuid)
}

func (rep *URLRepository) DeleteOlderRows(ctx context.Context, age time.Duration) error {
	defer rep.store.Commit(true)
	return rep.store.memStore.URL().DeleteOlderRows(ctx, age)
}

func (rep *URLRepository) getURLRepository() repository.URLRepository {
	return rep.store.memStore.URL()
}
