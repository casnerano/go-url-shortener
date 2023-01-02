package sqlstore

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

type Store struct {
	pgxpool *pgxpool.Pool
}

func NewStore(pgxpool *pgxpool.Pool) *Store {
	return &Store{pgxpool: pgxpool}
}

func (s *Store) URL() repository.URLRepository {
	return &URLRepository{store: s}
}
