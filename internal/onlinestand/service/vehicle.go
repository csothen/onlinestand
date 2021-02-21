package service

import (
	"io"

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

// GetByID : Retrieves a single instance of the model Vehicle based on an id
func (service *VehicleService) GetByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// GetAll : Retrieves all instances of the model Vehicle
func (service *VehicleService) GetAll() (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// GetAllAvailable : Retreives all instances of the model Vehicle where the status is AVAILABLE
func (service *VehicleService) GetAllAvailable() (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// Create : Creates and persists a new instance of the model Vehicle
func (service *VehicleService) Create(body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// UpdateByID : Updates and persists an instance of the model Vehicle that matches a given ID
func (service *VehicleService) UpdateByID(id string, body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// DeleteByID : Deletes an instance of the model Vehicle that matches a given ID
func (service *VehicleService) DeleteByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}
