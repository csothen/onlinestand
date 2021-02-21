package controller

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
	"github.com/csothen/onlinestand/pkg/responsehandler"
	"github.com/gorilla/mux"
)

type UserController struct {
	l       *log.Logger
	service *service.UserService
}

func NewUserController(l *log.Logger) *UserController {
	service := service.NewUserService()
	return &UserController{l, service}
}

func (uc *UserController) GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	uc.l.Println("GET Users")

	res, err := uc.service.GetAll()
	responsehandler.Respond(uc.l, rw, res, err)
}

func (uc *UserController) GetUserById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	uc.l.Println("GET User By ID", id)

	res, err := uc.service.GetByID(id)
	responsehandler.Respond(uc.l, rw, res, err)
}

func (uc *UserController) GetUserByUsername(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	uc.l.Println("GET User by Username", username)

	res, err := uc.service.GetByUsername(username)
	responsehandler.Respond(uc.l, rw, res, err)
}

func (uc *UserController) CreateUser(rw http.ResponseWriter, r *http.Request) {
	uc.l.Println("POST New User")

	res, err := uc.service.Create(r.Body)
	responsehandler.Respond(uc.l, rw, res, err)
}

func (uc *UserController) UpdateUserById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	uc.l.Println("PUT User by ID", id)

	res, err := uc.service.UpdateByID(id, r.Body)
	responsehandler.Respond(uc.l, rw, res, err)
}

func (uc *UserController) UpdateUserByUsername(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	uc.l.Println("PUT User by Username", username)

	res, err := uc.service.UpdateByUsername(username, r.Body)
	responsehandler.Respond(uc.l, rw, res, err)
}

func (uc *UserController) DeleteUserById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	uc.l.Println("DELETE User by ID", id)

	res, err := uc.service.DeleteByID(id)
	responsehandler.Respond(uc.l, rw, res, err)
}

func (uc *UserController) DeleteUserByUsername(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	uc.l.Println("DELETE User by Username", username)

	res, err := uc.service.DeleteByUsername(username)
	responsehandler.Respond(uc.l, rw, res, err)
}
