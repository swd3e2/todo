package controllers

import (
	"net/http"

	"github.com/swd3e2/todo/internal/container"
	"github.com/swd3e2/todo/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(cont *container.Container) *UserController {
	return &UserController{
		userService: cont.Get(container.USER_SERVICE).(*service.UserService),
	}
}

func (this *UserController) Register(w http.ResponseWriter, r *http.Request) {
	payload := service.RegisterPayload{}

	payload.Login = r.FormValue("login")
	payload.Password = r.FormValue("password")
	payload.Name = r.FormValue("name")
	payload.LastName = r.FormValue("lastname")

	if err := this.userService.Register(payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (this *UserController) Authorize(w http.ResponseWriter, r *http.Request) {
	if token, err := this.userService.Authorize(r.FormValue("login"), r.FormValue("password")); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token.Token))
	}
}
