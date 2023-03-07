package sqlstore

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/casnerano/go-url-shortener/internal/app/repository"
)

const (
	// MigrationSourceURL path to migration sources.
	MigrationSourceURL = "file://migrations/postgres"
)

// Store structure for sql store.
type Store struct {
	pgxpool *pgxpool.Pool
}

// NewStore constructor.
func NewStore(pgxpool *pgxpool.Pool) *Store {
	store := Store{pgxpool: pgxpool}
	// todo: logging
	_ = store.loadMigrations()
	return &store
}

func (s *Store) loadMigrations() error {
	m, err := migrate.New(MigrationSourceURL, s.pgxpool.Config().ConnString())
	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		// todo logging - no new migrations
		return nil
	}

	if err != nil {
		return err
	}

	// togo logging - migrations successfully loaded
	return nil
}

// URL return url repository with sql store.
func (s *Store) URL() repository.URLRepository {
	return &URLRepository{store: s}
}
