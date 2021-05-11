package handler

import (
	"github.com/swd3e2/todo/internal/application"
	"github.com/swd3e2/todo/internal/application/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorize(t *testing.T) {
	ts := httptest.NewServer(http.Handler(NewAuthorize(nil, application.NewUserService(mock.NewUserRepository()))))
	defer ts.Close()
}
