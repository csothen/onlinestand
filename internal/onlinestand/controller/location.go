package controller

import (
	"fmt"
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
	fmt.Fprint(rw, "GetAllLocations")
}

func (lc *LocationController) GetLocationById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "GetLocationById")
}

func (lc *LocationController) CreateLocation(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "CreateLocation")
}

func (lc *LocationController) UpdateLocationById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "UpdateLocationById")
}

func (lc *LocationController) DeleteLocationById(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "DeleteLocationById")
}
