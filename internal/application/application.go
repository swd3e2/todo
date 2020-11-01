package application

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/container"
	"github.com/swd3e2/todo/internal/infrastructure/mysql"
	"github.com/swd3e2/todo/internal/repository"
	"github.com/swd3e2/todo/internal/service"
)

type Application struct {
	router    *Router
	config    *Config
	logger    *logrus.Logger
	store     *Store
	container *container.Container
}

func New() *Application {
	return &Application{
		logger:    logrus.New(),
		store:     NewStore(),
		router:    NewRouter(),
		config:    NewConfig(),
		container: container.New(),
	}
}

func (this *Application) Configure(filename string) error {
	if err := this.config.SetUp(filename); err != nil {
		return err
	}

	if level, err := logrus.ParseLevel(this.config.LogLevel); err != nil {
		return err
	} else {
		this.logger.SetLevel(level)
		this.logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	this.logger.Info(this.config)

	if err := this.store.Connect(this.config); err != nil {
		return err
	}
	this.logger.Info("Successfully connected to database")

	this.initRepositories()
	this.initServices()

	if err := this.router.Configure(this.container); err != nil {
		return err
	}

	return nil
}

func (this *Application) Run() error {
	if err := http.ListenAndServe(fmt.Sprintf(":%s", this.config.Port), this.router.Router); err != nil {
		return err
	}

	return nil
}

func (this *Application) initRepositories() {
	this.container.Set(container.USER_REPOSITORY, mysql.NewMySqlRepository(this.store.Db))
}

func (this *Application) initServices() {
	this.container.Set(container.USER_SERVICE, service.NewUserService(
		this.container.Get(container.USER_REPOSITORY).(repository.UserRepository),
	))
}
