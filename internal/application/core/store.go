package core

import (
	"github.com/jackc/pgx"
)

// Store Хранилище
type Store struct {
	Db *pgx.Conn
}

// NewStore Создание нового хранилища
func NewStore() *Store {
	return &Store{}
}

// Connect Открытие коннекта к бд
func (s *Store) Connect(config *StoreConfig) error {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     config.Host,
		Port:     uint16(config.Port),
		User:     config.User,
		Password: config.Password,
		Database: config.Database,
	})

	if err != nil {
		return err
	}

	s.Db = conn

	return nil
}

// Close Закрытие коннекта к бд
func (s *Store) Close() error {
	return s.Db.Close()
}
