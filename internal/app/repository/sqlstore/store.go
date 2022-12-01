package sqlstore

import (
	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type Store struct {
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) URL() repository.URLRepository {
	return &URLRepository{store: s}
}
