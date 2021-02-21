package controller

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
	"github.com/csothen/onlinestand/pkg/responsehandler"
	"github.com/gorilla/mux"
)

type LocationController struct {
	l       *log.Logger
	service *service.LocationService
}

func NewLocationController(l *log.Logger) *LocationController {
	service := service.NewLocationService()
	return &LocationController{l, service}
}

func (lc *LocationController) GetAllLocations(rw http.ResponseWriter, r *http.Request) {
	lc.l.Println("GET Locations")

	res, err := lc.service.GetAll()
	responsehandler.Respond(lc.l, rw, res, err)
}

func (lc *LocationController) GetLocationById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	lc.l.Println("GET Location By ID", id)

	res, err := lc.service.GetByID(id)
	responsehandler.Respond(lc.l, rw, res, err)
}

func (lc *LocationController) CreateLocation(rw http.ResponseWriter, r *http.Request) {
	lc.l.Println("POST New Location")

	res, err := lc.service.Create(r.Body)
	responsehandler.Respond(lc.l, rw, res, err)
}

func (lc *LocationController) UpdateLocationById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	lc.l.Println("PUT Location By ID", id)

	res, err := lc.service.UpdateByID(id, r.Body)
	responsehandler.Respond(lc.l, rw, res, err)
}

func (lc *LocationController) DeleteLocationById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	lc.l.Println("DELETE Location By ID", id)

	res, err := lc.service.DeleteByID(id)
	responsehandler.Respond(lc.l, rw, res, err)
}
