package controller

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
	"github.com/csothen/onlinestand/pkg/responsehandler"
	"github.com/gorilla/mux"
)

type VehicleController struct {
	l       *log.Logger
	service *service.VehicleService
}

func NewVehicleController(l *log.Logger) *VehicleController {
	service := service.NewVehicleService()
	return &VehicleController{l, service}
}

func (vc *VehicleController) GetAllVehicles(rw http.ResponseWriter, r *http.Request) {
	vc.l.Println("GET Vehicles")

	res, err := vc.service.GetAll()
	responsehandler.Respond(vc.l, rw, res, err)
}

func (vc *VehicleController) GetVehicleById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vc.l.Println("GET Vehicle by ID", id)

	res, err := vc.service.GetByID(id)
	responsehandler.Respond(vc.l, rw, res, err)
}

func (vc *VehicleController) GetAllAvailableVehicles(rw http.ResponseWriter, r *http.Request) {
	vc.l.Println("GET Available Vehicles")

	res, err := vc.service.GetAllAvailable()
	responsehandler.Respond(vc.l, rw, res, err)
}

func (vc *VehicleController) CreateVehicle(rw http.ResponseWriter, r *http.Request) {
	vc.l.Println("POST New Vehicle")

	res, err := vc.service.Create(r.Body)
	responsehandler.Respond(vc.l, rw, res, err)
}

func (vc *VehicleController) UpdateVehicleById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vc.l.Println("PUT Vehicle by ID")

	res, err := vc.service.UpdateByID(id, r.Body)
	responsehandler.Respond(vc.l, rw, res, err)
}

func (vc *VehicleController) DeleteVehicleById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vc.l.Println("DELETE Vehicle by ID")

	res, err := vc.service.DeleteByID(id)
	responsehandler.Respond(vc.l, rw, res, err)
}
