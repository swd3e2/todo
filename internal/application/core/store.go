package core

import (
	"github.com/jackc/pgx"
)

// Store Хранилище
type Store struct {
	conn *pgx.Conn
}

// NewStore Создание нового хранилища
func NewStore() *Store {
	return &Store{}
}

// Connect Открытие коннекта к бд
func (s *Store) Connect(config *Config) error {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     config.Store.Host,
		Port:     uint16(config.Store.Port),
		User:     config.Store.User,
		Password: config.Store.Password,
		Database: config.Store.Database,
	})

	if err != nil {
		return err
	}

	s.conn = conn

	return nil
}

// Close Закрытие коннекта к бд
func (s *Store) Close() error {
	return s.conn.Close()
}
