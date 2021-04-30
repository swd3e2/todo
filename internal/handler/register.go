package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type RegisterHandler struct {
	logger *logrus.Logger
}

func NewRegister(l *logrus.Logger) *RegisterHandler {
	return &RegisterHandler{l}
}

func (h *RegisterHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
