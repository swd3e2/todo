package handler

import (
	"bytes"
	"github.com/bmizerany/assert"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
	"github.com/swd3e2/todo/internal/application"
	"github.com/swd3e2/todo/internal/application/generated_mocks"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler(t *testing.T) {
	type mockBehavior func(s *mock_application.MockUserService, user application.User)

	counter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests",
		}, []string{"path"})

	logger := logrus.New()
	logger.Out = ioutil.Discard

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            application.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test", "lastname": "Test", "login": "admin", "password": "admin"}`,
			inputUser: application.User{
				Name:     "Test",
				LastName: "Test",
				Login:    "admin",
				Password: "admin",
			},
			mockBehavior: func(s *mock_application.MockUserService, user application.User) {
				s.EXPECT().Register(user.Name, user.LastName, user.Login, user.Password).Return(nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: ``,
		},
		{
			name:      "Validation fail",
			inputBody: `{"lastname": "Test", "login": "admin", "password": "admin"}`,
			inputUser: application.User{
				Name:     "Test",
				LastName: "Test",
				Age:      0,
				Login:    "admin",
				Password: "admin",
			},
			mockBehavior:         func(s *mock_application.MockUserService, user application.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: "Bad Request\n",
		},
		{
			name:      "Login is already in use",
			inputBody: `{"name": "Test", "lastname": "Test", "login": "admin", "password": "admin"}`,
			inputUser: application.User{
				Name:     "Test",
				LastName: "Test",
				Age:      0,
				Login:    "admin",
				Password: "admin",
			},
			mockBehavior: func(s *mock_application.MockUserService, user application.User) {
				s.EXPECT().Register(user.Name, user.LastName, user.Login, user.Password).Return(application.LoginIsAlreadyInUse)
			},
			expectedStatusCode:   422,
			expectedResponseBody: "Unprocessable Entity\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			userService := mock_application.NewMockUserService(c)
			testCase.mockBehavior(userService, testCase.inputUser)

			router := mux.NewRouter()
			router.Handle("/register", NewRegister(logger, counter, userService)).Methods("POST")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(testCase.inputBody))
			router.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
