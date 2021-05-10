package core

import (
	"context"
	"fmt"
	"github.com/swd3e2/todo/internal/application"
	"github.com/swd3e2/todo/internal/application/handler"
	"github.com/swd3e2/todo/internal/application/postgres"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Application Корневая структура приложения
type Application struct {
	router *mux.Router
	config *Config
	logger *logrus.Logger
	store  *Store
}

// New Создание нового приложения
func New() *Application {
	return &Application{
		logger: logrus.New(),
		store:  NewStore(),
		config: NewConfig(),
	}
}

// Configure Загрузка настроек приложения
func (a *Application) Configure(filename string) error {
	if err := a.config.SetUp(filename); err != nil {
		return err
	}

	if level, err := logrus.ParseLevel(a.config.LogLevel); err != nil {
		return err
	} else {
		a.logger.SetLevel(level)
		a.logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	a.logger.Info(a.config)

	if err := a.store.Connect(a.config); err != nil {
		return err
	}
	a.logger.Info("Successfully connected to database")

	service := application.NewUserService(postgres.NewUserRepository(a.store.conn))

	a.router = mux.NewRouter()
	a.router.Handle("/user/register", handler.NewRegister(a.logger, service)).Methods("POST")
	a.router.Handle("/user/authorize", handler.NewAuthorize(a.logger, service)).Methods("POST")
	a.router.Handle("/todo", handler.NewCreateTodo(a.logger)).Methods("POST")

	return nil
}

// Run Запуск приложения
func (a *Application) Run() error {
	defer a.store.Close()

	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		a.config.Store.User,
		a.config.Store.Password,
		a.config.Store.Host,
		a.config.Store.Port,
		a.config.Store.Database,
	)

	if err := postgres.RunMigrations(dbUrl, a.config.MigrationsPath); err != nil {
		a.logger.WithError(err).Error("Не смогли применить миграцию")
		return err
	}
	a.logger.Info("Миграции применены")

	server := http.Server{
		Addr:         ":" + a.config.Port,
		Handler:      a.router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	doneChan := make(chan struct{})

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	go func() {
		server.ListenAndServe()
		doneChan <- struct{}{}
	}()

	go func() {
		<-sigChan
		doneChan <- struct{}{}
	}()

	<-doneChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	server.Shutdown(ctx)
	cancel()

	return nil
}