package service

import (
	"io"

	"github.com/csothen/onlinestand/internal/onlinestand/repository/postgres"
	"github.com/csothen/onlinestand/pkg/models"
)

// LocationService : Implementation of the interface LocationService
type LocationService struct {
	repo models.LocationRepository
}

// NewLocationService : Creates a new instance of LocationService that handles business logic related to the Location model
func NewLocationService() *LocationService {
	repo := postgres.NewLocationRepository()

	return &LocationService{repo}
}

// GetAll : Retrieves all instances of the model Location
func (service *LocationService) GetAll() (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// GetByID : Retrieves an instance of the model Location that matches a given ID
func (service *LocationService) GetByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// Create : Creates and persists a new instance of the model Location
func (service *LocationService) Create(body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// UpdateByID : Updates and persists an instance of the model Location that matches a given ID
func (service *LocationService) UpdateByID(id string, body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// DeleteByID : Deletes an instance of the model Location that matches a given ID
func (service *LocationService) DeleteByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}
