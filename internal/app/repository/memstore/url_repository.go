package memstore

import (
	"context"
	"errors"
	"time"

	"github.com/casnerano/go-url-shortener/internal/app/model"
)

type URLRepository struct {
	store   *Store
	counter int
}

func (rep *URLRepository) Add(_ context.Context, url *model.ShortURL) error {
	rep.counter++
	url.ID = rep.counter
	url.CreatedAt = time.Now()
	rep.store.ShortURLStorage[url.Code] = url
	return nil
}

func (rep *URLRepository) GetByCode(_ context.Context, code string) (*model.ShortURL, error) {
	url, ok := rep.store.ShortURLStorage[code]
	if !ok {
		return nil, errors.New("url not found")
	}
	return url, nil
}

func (rep *URLRepository) FindByUser(ctx context.Context, uid model.UserID) ([]*model.ShortURL, error) {
	collection := []*model.ShortURL{}
	for _, shortURL := range rep.store.ShortURLStorage {
		if shortURL.UserID == uid {
			collection = append(collection, shortURL)
		}
	}
	return collection, nil
}

func (rep *URLRepository) DeleteByCode(_ context.Context, code string) error {
	_, ok := rep.store.ShortURLStorage[code]
	if !ok {
		return errors.New("url not found")
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
