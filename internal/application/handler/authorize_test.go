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
	mock_application "github.com/swd3e2/todo/internal/application/generated_mocks"
	"net/http/httptest"
	"testing"
)

func TestAuthorize(t *testing.T) {
	type mockBehavior func(s *mock_application.MockUserService, user application.User)

	counter := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests",
		}, []string{"path"})

	logger := logrus.New()
	//logger.Out = ioutil.Discard

	testTable := []struct {
		name                 string
		inputBody            string
		user                 application.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"login": "admin", "password": "admin"}`,
			user: application.User{
				Login:    "admin",
				Password: "admin",
			},
			mockBehavior: func(s *mock_application.MockUserService, user application.User) {
				s.EXPECT().Authorize(user.Login, user.Password).Return(&application.Token{"some_token"}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "some_token",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_application.NewMockUserService(c)
			testCase.mockBehavior(service, testCase.user)

			router := mux.NewRouter()
			router.Handle("/authorize", NewAuthorize(logger, counter, service))

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/authorize", bytes.NewBufferString(testCase.inputBody))

			router.ServeHTTP(w, r)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
