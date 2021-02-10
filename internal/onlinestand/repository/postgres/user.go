package postgres

import (
	"database/sql"

	"github.com/csothen/onlinestand/pkg/models"
)

// UserRepository : Postgres Repository that handles transactions related to the User model
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository : Creates a new instance of UserRepository with an open connection to the Postgres database
func NewUserRepository() *UserRepository {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	return &UserRepository{db}
}

// CreateUser : Persists a new instance of the model User in the database
func (repo *UserRepository) CreateUser(c models.User) error {
	panic("Not implemented!")
}

// GetAllUser : Retrieves all instances of the model User in the database
func (repo *UserRepository) GetAllUser() ([]*models.User, error) {
	panic("Not implemented!")
}
