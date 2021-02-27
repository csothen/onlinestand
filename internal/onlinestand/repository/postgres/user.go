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
func (repo *UserRepository) CreateUser(user models.User) (*models.User, error) {
	panic("Not implemented!")
}

// GetAllUsers : Retrieves all instances of the model User in the database
func (repo *UserRepository) GetAllUsers() ([]*models.User, error) {
	panic("Not implemented!")
}

// GetUserByID : Retrieves an instance of user that matches a given ID
// or nil in case there is no match
func (repo *UserRepository) GetUserByID(id string) (*models.User, error) {
	panic("Not implemented!")
}

// GetUserByUsername : Retrieves an instance of user that matches a given username
// or nil in case there is no match
func (repo *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	panic("Not implemented!")
}

// DeleteUserByID : Deletes an instance of user that matches a given ID
func (repo *UserRepository) DeleteUserByID(id string) error {
	panic("Not implemented!")
}

// DeleteUserByUsername : Deletes an instance of user that matches a given username
func (repo *UserRepository) DeleteUserByUsername(username string) error {
	panic("Not implemented!")
}

// Save : Persists an existing instance of the model User in the database
func (repo *UserRepository) Save(user *models.User) (*models.User, error) {
	panic("Not implemented!")
}
