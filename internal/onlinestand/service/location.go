package service

import (
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

// CreateLocation : Creates and persists a new instance of the model Location
func (service *LocationService) CreateLocation(c models.Location) error {
	panic("Not implemented!")
}

// GetAllLocation : Retrieves all instances of the model Location
func (service *LocationService) GetAllLocation() ([]*models.Location, error) {
	panic("Not implemented!")
}
