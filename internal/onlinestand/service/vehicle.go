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
func (service *VehicleService) CreateVehicle(c models.Vehicle) error {
	panic("Not implemented!")
}

// GetAllVehicle : Retrieves all instances of the model Vehicle
func (service *VehicleService) GetAllVehicle() ([]*models.Vehicle, error) {
	panic("Not implemented!")
}
