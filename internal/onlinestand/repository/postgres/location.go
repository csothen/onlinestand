package postgres

import (
	"database/sql"

	"github.com/csothen/onlinestand/pkg/models"
)

// LocationRepository : Postgres Repository that handles transactions related to the Location model
type LocationRepository struct {
	db *sql.DB
}

// NewLocationRepository : Creates a new instance of LocationRepository with an open connection to the Postgres database
func NewLocationRepository() *LocationRepository {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	return &LocationRepository{db}
}

// CreateLocation : Persists a new instance of the model Location in the database
func (repo *LocationRepository) CreateLocation(c models.Location) error {
	panic("Not implemented!")
}

// GetAllLocation : Retrieves all instances of the model Location in the database
func (repo *LocationRepository) GetAllLocation() ([]*models.Location, error) {
	panic("Not implemented!")
}
