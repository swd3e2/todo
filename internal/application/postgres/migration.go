package postgres

import (
	"errors"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

var (
	noDbProvided            = errors.New("не переданы реквизиты подключения к бд")
	noMigrationPathProvided = errors.New("не передан путь миграции")
)

// RunMigrations Применяет миграции к бд
func RunMigrations(db string, path string) error {
	if db == "" {
		return noDbProvided
	}

	if path == "" {
		return noMigrationPathProvided
	}

	m, err := migrate.New(path, db)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
