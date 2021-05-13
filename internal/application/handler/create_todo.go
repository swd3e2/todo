package handler

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CreateTodoHandler struct {
	logger  *logrus.Logger
	counter *prometheus.CounterVec
}

func NewCreateTodo(logger *logrus.Logger) *CreateTodoHandler {
	return &CreateTodoHandler{
		logger: logger,
	}
}

func (h *CreateTodoHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CreateTodoHandler")
}
