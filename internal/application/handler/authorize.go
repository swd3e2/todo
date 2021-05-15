package handler

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/application"
	"net/http"
)

// swagger:route POST /authorize auth authorize
// Authorize
//
// responses:
//	201: noContentResponse
//  400: errorResponse
//  500: errorResponse

type AuthorizeHandler struct {
	logger  *logrus.Logger
	counter *prometheus.CounterVec
	service application.UserService
}

func NewAuthorize(
	logger *logrus.Logger,
	counter *prometheus.CounterVec,
	service application.UserService,
) *AuthorizeHandler {
	return &AuthorizeHandler{
		logger:  logger,
		counter: counter,
		service: service,
	}
}

func (h *AuthorizeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Infof("AuthorizeHandler")
	h.counter.With(prometheus.Labels{"path": "/authorize"}).Inc()

	user := struct {
		Login    string `validate:"required" json:"login"`
		Password string `validate:"required" json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.logger.Error("Не удалось распарсить запрос")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		h.logger.Error("Не передан логин или пароль")
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	token, err := h.service.Authorize(user.Login, user.Password)
	if err != nil {
		h.logger.WithError(err).Error("Ошибка в севисе")
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, _ = rw.Write([]byte(token.Token))
}
