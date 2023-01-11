package memstore

import (
	"context"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type URLRepository struct {
	store   *Store
	counter int
}

func (rep *URLRepository) Add(_ context.Context, url *model.ShortURL) error {
	for _, shortURL := range rep.store.ShortURLStorage {
		if shortURL.UserUUID == url.UserUUID && shortURL.Original == url.Original {
			return repository.ErrURLExist
		}
	}

	rep.counter++
	url.ID = rep.counter
	url.CreatedAt = time.Now()
	rep.store.ShortURLStorage[url.Code] = url

	return nil
}

func (rep *URLRepository) AddBatch(ctx context.Context, urls []*model.ShortURL) error {
	for _, url := range urls {
		_ = rep.Add(ctx, url)
	}
	return nil
}

func (rep *URLRepository) GetByCode(_ context.Context, code string) (*model.ShortURL, error) {
	url, ok := rep.store.ShortURLStorage[code]
	if !ok {
		return nil, repository.ErrURLNotFound
	}
	return url, nil
}

func (rep *URLRepository) GetByUserUUIDAndOriginal(_ context.Context, uuid string, original string) (*model.ShortURL, error) {
	for _, shortURL := range rep.store.ShortURLStorage {
		if shortURL.UserUUID == uuid && shortURL.Original == original {
			return shortURL, nil
		}
	}
	return nil, repository.ErrURLNotFound
}

func (rep *URLRepository) FindByUserUUID(ctx context.Context, uuid string) ([]*model.ShortURL, error) {
	collection := []*model.ShortURL{}
	for _, shortURL := range rep.store.ShortURLStorage {
		if shortURL.UserUUID == uuid {
			collection = append(collection, shortURL)
		}
	}
	return collection, nil
}

func (rep *URLRepository) DeleteByCode(_ context.Context, code string) error {
	_, ok := rep.store.ShortURLStorage[code]
	if !ok {
		return repository.ErrURLNotFound
	}

	delete(rep.store.ShortURLStorage, code)
	return nil
}

func (rep *URLRepository) DeleteOlderRows(_ context.Context, d time.Duration) error {
	for code, shortURL := range rep.store.ShortURLStorage {
		if shortURL.CreatedAt.Add(d).Before(time.Now()) {
			delete(rep.store.ShortURLStorage, code)
		}
	}
	return nil
}
