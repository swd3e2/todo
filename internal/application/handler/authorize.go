package handler

import (
	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/application"
	"net/http"
)

type AuthorizeHandler struct {
	logger  *logrus.Logger
	service *application.UserService
}

func NewAuthorize(logger *logrus.Logger, service *application.UserService) *AuthorizeHandler {
	return &AuthorizeHandler{
		logger:  logger,
		service: service,
	}
}

func (h *AuthorizeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Infof("AuthorizeHandler")

	if err := r.ParseForm(); err != nil {
		h.logger.Error("Не удалось распаристь форму")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	login := r.Form.Get("login")
	password := r.Form.Get("password")

	if login == "" || password == "" {
		h.logger.Error("Не передан логин или пароль")
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	t, err := h.service.Authorize(login, password)
	if err != nil {
		h.logger.WithError(err).Error("Ошибка в севисе")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, _ = rw.Write([]byte(t.Token))
}
