package memstore

import (
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

// ShortURLStorage alias type
type ShortURLStorage = map[string]*model.ShortURL

// Store structure for memory store.
type Store struct {
	ShortURLStorage ShortURLStorage
}

// NewStore constructor.
func NewStore() *Store {
	return &Store{
		ShortURLStorage: make(ShortURLStorage),
	}
}

// URL return url repository with memory store.
func (s *Store) URL() repository.URLRepository {
	return &URLRepository{store: s}
}
