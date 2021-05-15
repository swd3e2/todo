package test

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/application/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func (s *APITestSuite) TestRegister() {
	r := s.Require()

	logger := logrus.New()
	logger.Out = ioutil.Discard

	counter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests",
		}, []string{"path"})

	router := mux.NewRouter()
	router.Handle("/register", handler.NewRegister(logger, counter, s.userService))

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(`{"name": "Test", "lastname": "Test", "login": "admin", "password": "admin"}`))

	router.ServeHTTP(w, req)

	r.Equal(w.Code, http.StatusCreated)
}
