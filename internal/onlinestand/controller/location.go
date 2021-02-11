package controller

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
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

}

func (lc *LocationController) GetLocationById(rw http.ResponseWriter, r *http.Request) {

}

func (lc *LocationController) CreateLocation(rw http.ResponseWriter, r *http.Request) {

}

func (lc *LocationController) UpdateLocationById(rw http.ResponseWriter, r *http.Request) {

}

func (lc *LocationController) DeleteLocationById(rw http.ResponseWriter, r *http.Request) {

}
