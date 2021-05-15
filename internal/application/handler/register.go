package handler

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/application"
	"net/http"
)

type RegisterHandler struct {
	logger  *logrus.Logger
	counter *prometheus.CounterVec
	service application.UserService
}

func NewRegister(
	logger *logrus.Logger,
	counter *prometheus.CounterVec,
	service application.UserService,
) *RegisterHandler {
	return &RegisterHandler{
		logger:  logger,
		counter: counter,
		service: service,
	}
}

func (h *RegisterHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Infof("RegisterHandler")
	h.counter.With(prometheus.Labels{"path": "/register"}).Inc()

	user := struct {
		Name     string `validate:"required" json:"name"`
		Lastname string `validate:"required" json:"lastname"`
		Login    string `validate:"required" json:"login"`
		Password string `validate:"required" json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.WithError(err).Error("Не удалось распарсить запрос")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		h.logger.WithError(err).Error("Переданы не все параметры")
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.service.Register(user.Name, user.Lastname, user.Login, user.Password)
	if err == application.LoginIsAlreadyInUse {
		h.logger.WithError(err).Error("Логин занят")
		http.Error(rw, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	} else if err != nil {
		h.logger.WithError(err).Error("Ошибка в сервисе")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
