package memstore

import (
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type shortURLStorage = map[string]model.ShortURL

type Store struct {
	shortURLStorage shortURLStorage
}

func NewStore() *Store {
	return &Store{
		shortURLStorage: make(shortURLStorage),
	}
}

func (s *Store) URL() repository.URLRepository {
	return &URLRepository{store: s}
}
