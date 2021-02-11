package controller

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
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

}

func (uc *UserController) GetUserById(rw http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) GetUserByUsername(rw http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) CreateUser(rw http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) UpdateUserById(rw http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) UpdateUserByUsername(rw http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) DeleteUserById(rw http.ResponseWriter, r *http.Request) {

}

func (uc *UserController) DeleteUserByUsername(rw http.ResponseWriter, r *http.Request) {

}
