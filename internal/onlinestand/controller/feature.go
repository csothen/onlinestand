package controller

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
	"github.com/csothen/onlinestand/pkg/responsehandler"
	"github.com/gorilla/mux"
)

type FeatureController struct {
	l       *log.Logger
	service *service.FeatureService
}

func NewFeatureController(l *log.Logger) *FeatureController {
	service := service.NewFeatureService()
	return &FeatureController{l, service}
}

func (fc *FeatureController) GetAllFeatures(rw http.ResponseWriter, r *http.Request) {
	fc.l.Println("GET Features")

	res, err := fc.service.GetAll()
	responsehandler.Respond(fc.l, rw, res, err)
}

func (fc *FeatureController) GetFeatureById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fc.l.Println("GET Feature By ID", id)

	res, err := fc.service.GetByID(id)
	responsehandler.Respond(fc.l, rw, res, err)
}

func (fc *FeatureController) CreateFeature(rw http.ResponseWriter, r *http.Request) {
	fc.l.Println("POST New Feature")

	res, err := fc.service.Create(r.Body)
	responsehandler.Respond(fc.l, rw, res, err)
}

func (fc *FeatureController) UpdateFeatureById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fc.l.Println("PUT Feature By ID", id)

	res, err := fc.service.UpdateByID(id, r.Body)
	responsehandler.Respond(fc.l, rw, res, err)
}

func (fc *FeatureController) DeleteFeatureById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fc.l.Println("DELETE Feature By ID", id)

	res, err := fc.service.DeleteByID(id)
	responsehandler.Respond(fc.l, rw, res, err)
}
