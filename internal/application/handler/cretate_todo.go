package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type CreateTodoHandler struct {
	logger *logrus.Logger
}

func NewCreateTodo(l *logrus.Logger) *CreateTodoHandler {
	return &CreateTodoHandler{l}
}

func (h *CreateTodoHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CreateTodoHandler")
}
