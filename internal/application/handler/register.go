package handler

import (
	"github.com/go-playground/validator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/application"
	"net/http"
)

type RegisterHandler struct {
	logger  *logrus.Logger
	counter *prometheus.CounterVec
	service *application.UserService
}

func NewRegister(
	logger *logrus.Logger,
	counter *prometheus.CounterVec,
	service *application.UserService,
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

	if err := r.ParseForm(); err != nil {
		h.logger.WithError(err).Error("Не удалось распарсить форму")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user := struct {
		Name     string `validate:"required"`
		Lastname string `validate:"required"`
		Login    string `validate:"required"`
		Password string `validate:"required"`
	}{
		Name:     r.Form.Get("name"),
		Lastname: r.Form.Get("lastname"),
		Login:    r.Form.Get("login"),
		Password: r.Form.Get("password"),
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		h.logger.WithError(err).Error("Переданы не все параметры")
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := h.service.Register(user.Name, user.Lastname, user.Login, user.Password); err != nil {
		h.logger.WithError(err).Error("Ошибка в севисе")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
