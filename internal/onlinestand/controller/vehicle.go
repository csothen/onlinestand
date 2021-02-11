package controller

import (
	"log"
	"net/http"

	"github.com/csothen/onlinestand/internal/onlinestand/service"
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

}

func (vc *VehicleController) GetVehicleById(rw http.ResponseWriter, r *http.Request) {

}

func (vc *VehicleController) GetAllAvailableVehicles(rw http.ResponseWriter, r *http.Request) {

}

func (vc *VehicleController) CreateVehicle(rw http.ResponseWriter, r *http.Request) {

}

func (vc *VehicleController) UpdateVehicleById(rw http.ResponseWriter, r *http.Request) {

}

func (vc *VehicleController) DeleteVehicleById(rw http.ResponseWriter, r *http.Request) {

}
