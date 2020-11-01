package application

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swd3e2/todo/internal/container"
	"github.com/swd3e2/todo/internal/controllers"
)

type Router struct {
	Router         *mux.Router
	userController *controllers.UserController
}

func NewRouter() *Router {
	return &Router{}
}

func (this *Router) Configure(container *container.Container) error {
	this.userController = controllers.NewUserController(container)

	this.Router = mux.NewRouter()

	this.Router.HandleFunc("/user/register", func(w http.ResponseWriter, r *http.Request) {
		this.userController.Register(w, r)
	}).Methods("POST")

	this.Router.HandleFunc("/user/authorize", func(w http.ResponseWriter, r *http.Request) {
		this.userController.Authorize(w, r)
	}).Methods("POST")

	return nil
}
