package memstore

import (
	"github.com/casnerano/go-url-shortener/internal/app/model"
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type ShortURLStorage = map[string]model.ShortURL

type Store struct {
	ShortURLStorage ShortURLStorage
}

func NewStore() *Store {
	return &Store{
		ShortURLStorage: make(ShortURLStorage),
	}
}

func (s *Store) URL() repository.URLRepository {
	return &URLRepository{store: s}
}
