package postgres

import (
	"database/sql"

	"github.com/csothen/onlinestand/pkg/models"
)

// VehicleRepository : Postgres Repository that handles transactions related to the Vehicle model
type VehicleRepository struct {
	db *sql.DB
}

// NewVehicleRepository : Creates a new instance of VehicleRepository with an open connection to the Postgres database
func NewVehicleRepository() *VehicleRepository {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	return &VehicleRepository{db}
}

// CreateVehicle : Persists a new instance of the model Vehicle in the database
func (repo *VehicleRepository) CreateVehicle(c models.Vehicle) error {
	panic("Not implemented!")
}

// GetVehicle : Retrieves a single instance of the model Vehicle from the database based on an id
func (repo *VehicleRepository) GetVehicle(id int) (*models.Vehicle, error) {
	panic("Not implemented!")
}

// GetAllVehicle : Retrieves all instances of the model Vehicle in the database
func (repo *VehicleRepository) GetAllVehicle() ([]*models.Vehicle, error) {
	panic("Not implemented!")
}

// GetAllAvailableVehicle : Retreives all instances of the model Vehicle from the database where the status is AVAILABLE
func (repo *VehicleRepository) GetAllAvailableVehicle() ([]*models.Vehicle, error) {
	panic("Not implemented!")
}
