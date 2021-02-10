package service

import (
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

// CreateUser : Creates and persists a new instance of the model User
func (service *UserService) CreateUser(c models.User) error {
	panic("Not implemented!")
}

// GetAllUser : Retrieves all instances of the model User
func (service *UserService) GetAllUser() ([]*models.User, error) {
	panic("Not implemented!")
}
