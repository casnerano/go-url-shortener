package migration

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

type MigrationDBType = string

const (
	MigrationDBTypePostgres MigrationDBType = "postgres"
)

const (
	MigrationSourcePath = "file://migrations"
)

type Migrator struct {
	dbDSN  string
	dbType MigrationDBType
}

func NewMigrator(dbType MigrationDBType, dbDSN string) *Migrator {
	return &Migrator{dbType: dbType, dbDSN: dbDSN}
}

func (migrator *Migrator) LoadMigrations() error {
	path := fmt.Sprintf("%s/%s", MigrationSourcePath, migrator.dbType)
	m, err := migrate.New(path, migrator.dbDSN)
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
