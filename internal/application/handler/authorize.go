package handler

import (
	"github.com/go-playground/validator"
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
	user := struct {
		Login    string `validate:"required"`
		Password string `validate:"required"`
	}{
		Login:    r.Form.Get("login"),
		Password: r.Form.Get("password"),
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		h.logger.Error("Не передан логин или пароль")
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	t, err := h.service.Authorize(user.Login, user.Password)
	if err != nil {
		h.logger.WithError(err).Error("Ошибка в севисе")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, _ = rw.Write([]byte(t.Token))
}