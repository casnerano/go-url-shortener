package memstore

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

// URLRepository structure for url repository with memory store.
type URLRepository struct {
	store   *Store
	counter int
}

// Adding entity.
func (rep *URLRepository) Add(_ context.Context, url *model.ShortURL) error {
	for _, shortURL := range rep.store.ShortURLStorage {
		if shortURL.UserUUID == url.UserUUID && shortURL.Original == url.Original {
			return repository.ErrURLAlreadyExist
		}
	}

	rep.counter++
	url.ID = rep.counter
	url.CreatedAt = time.Now()
	rep.store.ShortURLStorage[url.Code] = url

	return nil
}

// Batch adding entities.
func (rep *URLRepository) AddBatch(ctx context.Context, urls []*model.ShortURL) error {
	for _, url := range urls {
		_ = rep.Add(ctx, url)
	}
	return nil
}

// Get entity by short code.
func (rep *URLRepository) GetByCode(_ context.Context, code string) (*model.ShortURL, error) {
	url, ok := rep.store.ShortURLStorage[code]
	if !ok {
		return nil, repository.ErrURLNotFound
	}

	if url.Deleted {
		return nil, repository.ErrURLMarkedForDelete
	}

	return url, nil
}

// Get entity by user uuid and original url.
func (rep *URLRepository) GetByUserUUIDAndOriginal(_ context.Context, uuid string, original string) (*model.ShortURL, error) {
	for _, shortURL := range rep.store.ShortURLStorage {
		if shortURL.UserUUID == uuid && shortURL.Original == original {
			if shortURL.Deleted {
				return nil, repository.ErrURLMarkedForDelete
			}
			return shortURL, nil
		}
	}
	return nil, repository.ErrURLNotFound
}

// Find entities by user uuid.
func (rep *URLRepository) FindByUserUUID(_ context.Context, uuid string) ([]*model.ShortURL, error) {
	collection := []*model.ShortURL{}
	for _, shortURL := range rep.store.ShortURLStorage {
		if shortURL.UserUUID == uuid && !shortURL.Deleted {
			collection = append(collection, shortURL)
		}
	}
	return collection, nil
}

// Delete entity by short code.
func (rep *URLRepository) DeleteByCode(_ context.Context, code string, uuid string) error {
	_, ok := rep.store.ShortURLStorage[code]
	if !ok {
		return repository.ErrURLNotFound
	}

	if rep.store.ShortURLStorage[code].UserUUID == uuid {
		rep.store.ShortURLStorage[code].Deleted = true
		return nil
	}

	return repository.ErrURLNotFound
}

// Batch delete entities by codes.
func (rep *URLRepository) DeleteBatchByCodes(ctx context.Context, codes []string, uuid string) error {
	for _, code := range codes {
		rep.DeleteByCode(ctx, code, uuid)
	}
	return nil
}

// Delete older entities for duration.
func (rep *URLRepository) DeleteOlderRows(_ context.Context, d time.Duration) error {
	for code, shortURL := range rep.store.ShortURLStorage {
		if shortURL.CreatedAt.Add(d).Before(time.Now()) {
			rep.store.ShortURLStorage[code].Deleted = true
		}
	}
	return nil
}
