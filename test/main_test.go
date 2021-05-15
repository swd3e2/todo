package test

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/stretchr/testify/suite"
	"github.com/swd3e2/todo/internal/application"
	"github.com/swd3e2/todo/internal/application/core"
	"github.com/swd3e2/todo/internal/application/postgres"
	"testing"
)

type APITestSuite struct {
	suite.Suite
	db *pgx.Conn

	userService application.UserService
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(APITestSuite))
}

func (s *APITestSuite) SetupSuite() {
	storeConfig := &core.StoreConfig{
		Host:     "localhost",
		Port:     5433,
		User:     "admin",
		Password: "admin",
		Database: "test",
	}

	store := core.NewStore()
	if err := store.Connect(storeConfig); err != nil {
		s.FailNow("Не удалось подключится к бд", err)
	}
	s.db = store.Db

	s.initDeps()

	if err := s.populateDB(storeConfig); err != nil {
		s.FailNow("Failed to populate DB", err)
	}
}

func (s *APITestSuite) TearDownSuite() {
	_ = s.db.Close()
}

func (s *APITestSuite) initDeps() {
	userRepository := postgres.NewUserRepository(s.db)
	s.userService = application.NewUserService(userRepository)
}

func (s *APITestSuite) populateDB(config *core.StoreConfig) error {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	if err := postgres.RunMigrations(dbUrl, "file://../internal/application/postgres/migrations"); err != nil {
		return err
	}

	if _, err := s.db.Exec("truncate table users;"); err != nil {
		return err
	}

	//if _, err := s.db.Exec("truncate table todos;"); err != nil {
	//	return err
	//}

	return nil
}
