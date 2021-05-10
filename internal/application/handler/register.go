package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/application"
	"net/http"
)

type RegisterHandler struct {
	logger  *logrus.Logger
	service *application.UserService
}

func NewRegister(logger *logrus.Logger, service *application.UserService) *RegisterHandler {
	return &RegisterHandler{
		logger:  logger,
		service: service,
	}
}

func (h *RegisterHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Infof("RegisterHandler")

	if err := r.ParseForm(); err != nil {
		h.logger.WithError(err).Error("Не удалось распарсить форму")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	name := r.Form.Get("name")
	lastname := r.Form.Get("lastname")
	login := r.Form.Get("login")
	password := r.Form.Get("password")

	if login == "" || password == "" {
		h.logger.Error("Переданы не параметры")
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.service.Register(name, lastname, login, password)
	if err != nil {
		h.logger.WithError(err).Error("Ошибка в севисе")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
