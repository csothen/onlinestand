package service

import (
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

// CreateFeature : Creates and persists a new instance of the model Feature
func (service *FeatureService) CreateFeature(c models.Feature) error {
	panic("Not implemented!")
}

// GetAllFeature : Retrieves all instances of the model Feature
func (service *FeatureService) GetAllFeature() ([]*models.Feature, error) {
	panic("Not implemented!")
}
