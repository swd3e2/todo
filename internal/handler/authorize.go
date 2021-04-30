package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthorizeHandler struct {
	logger *logrus.Logger
}

func NewAuthorize(l *logrus.Logger) *AuthorizeHandler {
	return &AuthorizeHandler{l}
}

func (h *AuthorizeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
