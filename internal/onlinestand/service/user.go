package service

import (
	"io"

	"github.com/csothen/onlinestand/internal/onlinestand/repository/postgres"
	"github.com/csothen/onlinestand/pkg/models"
)

// UserService : Implementation of the interface UserService
type UserService struct {
	repo models.UserRepository
}

// NewUserService : Creates a new instance of UserService that handles business logic related to the User model
func NewUserService() *UserService {
	repo := postgres.NewUserRepository()

	return &UserService{repo}
}

// GetAll : Retrieves all instances of the model User
func (service *UserService) GetAll() (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// GetByID : Retrieves an instance of the model User that matches a given ID
func (service *UserService) GetByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// GetByUsername : Retrieves an instance of the model User that matches a given username
func (service *UserService) GetByUsername(username string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// Create : Creates and persists a new instance of the model User
func (service *UserService) Create(body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// UpdateByID : Updates an instance of User that matches a given id
func (service *UserService) UpdateByID(id string, body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// UpdateByUsername : Updates an instance of User that matches a given username
func (service *UserService) UpdateByUsername(username string, body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// DeleteByID : Deletes an instance of User that matches a given id
func (service *UserService) DeleteByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// DeleteByUsername : Deletes an instance of User that matches a given username
func (service *UserService) DeleteByUsername(username string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}
