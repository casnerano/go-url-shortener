package migration

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	"github.com/casnerano/go-url-shortener/internal/app/config"
)

const (
	MigrationSourceURL = "file://migrations"
)

type Manager struct {
	cfg *config.Config
}

func NewManager(c *config.Config) *Manager {
	return &Manager{cfg: c}
}

func (manager *Manager) LoadMigrations() error {
	m, err := migrate.New(MigrationSourceURL, manager.cfg.Storage.DSN)
	if err != nil {
		return err
	}

	err = m.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		fmt.Println("no new migrations")
		return nil
	}

	if err != nil {
		return err
	}

	fmt.Println("Migrations successfully loaded")
	return nil
}
