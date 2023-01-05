package sqlstore

import (
	"github.com/jackc/pgx/v5"

	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) *Store {
	return &Store{db: db}
}

func (s *Store) URL() repository.URLRepository {
	return &URLRepository{store: s}
}
