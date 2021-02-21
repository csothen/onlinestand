package service

import (
	"io"

	"github.com/csothen/onlinestand/internal/onlinestand/repository/postgres"
	"github.com/csothen/onlinestand/pkg/models"
)

// FeatureService : Implementation of the interface FeatureService
type FeatureService struct {
	repo models.FeatureRepository
}

// NewFeatureService : Creates a new instance of FeatureService that handles business logic related to the Feature model
func NewFeatureService() *FeatureService {
	repo := postgres.NewFeatureRepository()

	return &FeatureService{repo}
}

// GetAll : Retrieves all instances of the model Feature
func (service *FeatureService) GetAll() (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// GetByID : Retrieves an instance of the model Feature that matches a given an ID
func (service *FeatureService) GetByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// Create : Creates and persists a new instance of the model Feature
func (service *FeatureService) Create(body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// UpdateByID : Updates and persists an instance of the model Feature that matches a given ID
func (service *FeatureService) UpdateByID(id string, body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// DeleteByID : Deletes an instance of the model Feature that matches a given ID
func (service *FeatureService) DeleteByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}
