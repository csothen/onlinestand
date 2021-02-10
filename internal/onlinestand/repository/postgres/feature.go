package postgres

import (
	"database/sql"

	"github.com/csothen/onlinestand/pkg/models"
)

// FeatureRepository : Postgres Repository that handles transactions related to the Feature model
type FeatureRepository struct {
	db *sql.DB
}

// NewFeatureRepository : Creates a new instance of FeatureRepository with an open connection to the Postgres database
func NewFeatureRepository() *FeatureRepository {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	return &FeatureRepository{db}
}

// CreateFeature : Persists a new instance of the model Feature in the database
func (repo *FeatureRepository) CreateFeature(c models.Feature) error {
	panic("Not implemented!")
}

// GetAllFeature : Retrieves all instances of the model Feature in the database
func (repo *FeatureRepository) GetAllFeature() ([]*models.Feature, error) {
	panic("Not implemented!")
}
