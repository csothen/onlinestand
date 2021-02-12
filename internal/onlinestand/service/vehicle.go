package service

import (
	"github.com/csothen/onlinestand/internal/onlinestand/repository/postgres"
	"github.com/csothen/onlinestand/pkg/models"
)

// VehicleService : Implementation of the interface VehicleService
type VehicleService struct {
	repo models.VehicleRepository
}

// NewVehicleService : Creates a new instance of VehicleService that handles business logic related to the Vehicle model
func NewVehicleService() *VehicleService {
	repo := postgres.NewVehicleRepository()

	return &VehicleService{repo}
}

// CreateVehicle : Creates and persists a new instance of the model Vehicle
func (service *VehicleService) CreateVehicle(c models.Vehicle) models.ServiceResponse {
	panic("Not implemented!")
}

// GetVehicle : Retrieves a single instance of the model Vehicle based on an id
func (service *VehicleService) GetVehicle(id int) models.ServiceResponse {
	panic("Not implemented!")
}

// GetAllVehicle : Retrieves all instances of the model Vehicle
func (service *VehicleService) GetAllVehicle() models.ServiceResponse {
	panic("Not implemented!")
}

// GetAllAvailableVehicle : Retreives all instances of the model Vehicle where the status is AVAILABLE
func (service *VehicleService) GetAllAvailableVehicle() models.ServiceResponse {
	panic("Not implemented!")
}
