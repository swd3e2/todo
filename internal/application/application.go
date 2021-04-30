package application

import (
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go.net/context"
	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/handler"
)

type Application struct {
	router *mux.Router
	config *Config
	logger *logrus.Logger
	store  *Store
}

func New() *Application {
	return &Application{
		logger: logrus.New(),
		store:  NewStore(),
		config: NewConfig(),
	}
}

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

	a.router = mux.NewRouter()
	a.router.Handle("/user/register", handler.NewRegister(a.logger)).Methods("POST")
	a.router.Handle("/user/authorize", handler.NewAuthorize(a.logger)).Methods("POST")

	return nil
}

func (a *Application) Run() error {
	server := http.Server{
		Addr:         ":" + a.config.Port,
		Handler:      a.router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	go func() {
		server.ListenAndServe()
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	server.Shutdown(ctx)
	a.store.Close(ctx)
	cancel()

	return nil
}
