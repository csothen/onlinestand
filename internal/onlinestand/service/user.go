package service

import (
	"fmt"
	"io"

	"github.com/csothen/onlinestand/internal/onlinestand/repository/postgres"
	"github.com/csothen/onlinestand/pkg/decode"
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
	users, err := service.repo.GetAllUsers()

	if err != nil {
		return models.ServiceResponse{Value: nil, Err: "Failed to retrieve the users", Status: 500}, err
	}

	return models.ServiceResponse{Value: users, Status: 200}, nil
}

// GetByID : Retrieves an instance of the model User that matches a given ID
func (service *UserService) GetByID(id string) (models.ServiceResponse, error) {
	user, err := service.repo.GetUserByID(id)

	if err != nil {
		return models.ServiceResponse{Value: nil, Err: fmt.Sprintf("Failed to find the user with ID %s", id), Status: 500}, err
	}

	if user == nil {
		return models.ServiceResponse{Value: nil, Err: fmt.Sprintf("User with ID %s not found", id), Status: 404}, nil
	}

	return models.ServiceResponse{Value: user, Status: 200}, nil
}

// GetByUsername : Retrieves an instance of the model User that matches a given username
func (service *UserService) GetByUsername(username string) (models.ServiceResponse, error) {
	user, err := service.repo.GetUserByUsername(username)

	if err != nil {
		return models.ServiceResponse{Err: fmt.Sprintf("Failed to find the user with username %s", username), Status: 500}, err
	}

	if user == nil {
		return models.ServiceResponse{Err: fmt.Sprintf("User with username %s not found", username), Status: 404}, nil
	}

	return models.ServiceResponse{Value: user, Status: 200}, nil
}

// Create : Creates and persists a new instance of the model User
func (service *UserService) Create(body io.ReadCloser) (models.ServiceResponse, error) {
	user := new(models.User)

	if err := decode.FromJSON(user, body); err != nil {
		return models.ServiceResponse{Err: "Unable to parse request body", Status: 400}, err
	}

	user, err := service.repo.CreateUser(*user)

	if err != nil {
		return models.ServiceResponse{Err: fmt.Sprintf("Failed to create the new user with username %s", user.Username), Status: 500}, err
	}

	return models.ServiceResponse{Value: user, Status: 201}, nil
}

// UpdateByID : Updates an instance of User that matches a given id
func (service *UserService) UpdateByID(id string, body io.ReadCloser) (models.ServiceResponse, error) {
	user := new(models.User)

	if err := decode.FromJSON(user, body); err != nil {
		return models.ServiceResponse{Err: "Unable to parse request body", Status: 400}, err
	}

	response, err := service.GetByID(id)

	if err != nil || response.Value == nil {
		return response, err
	}

	updatedUser, ok := response.Value.(*models.User)

	if !ok {
		return models.ServiceResponse{Err: "Something went wrong", Status: 500}, err
	}

	updatedUser.FirstName = user.FirstName
	updatedUser.LastName = user.LastName

	updatedUser, err = service.repo.Save(updatedUser)

	if err != nil {
		return models.ServiceResponse{Err: fmt.Sprintf("Failed to update user with ID %s", id), Status: 500}, err
	}

	return models.ServiceResponse{Value: updatedUser, Status: 200}, nil
}

// UpdateByUsername : Updates an instance of User that matches a given username
func (service *UserService) UpdateByUsername(username string, body io.ReadCloser) (models.ServiceResponse, error) {
	user := new(models.User)

	if err := decode.FromJSON(user, body); err != nil {
		return models.ServiceResponse{Err: "Unable to parse request body", Status: 400}, err
	}

	response, err := service.GetByUsername(username)

	if err != nil || response.Value == nil {
		return response, err
	}

	updatedUser, ok := response.Value.(*models.User)

	if !ok {
		return models.ServiceResponse{Err: "Something went wrong", Status: 500}, err
	}

	updatedUser.FirstName = user.FirstName
	updatedUser.LastName = user.LastName

	updatedUser, err = service.repo.Save(updatedUser)

	if err != nil {
		return models.ServiceResponse{Err: fmt.Sprintf("Failed to update user with username %s", username), Status: 500}, err
	}

	return models.ServiceResponse{Value: updatedUser, Status: 200}, nil
}

// DeleteByID : Deletes an instance of User that matches a given id
func (service *UserService) DeleteByID(id string) (models.ServiceResponse, error) {
	if err := service.repo.DeleteUserByID(id); err != nil {
		return models.ServiceResponse{Err: fmt.Sprintf("Failed to delete user with ID %s", id), Status: 500}, err
	}

	return models.ServiceResponse{Status: 200}, nil
}

// DeleteByUsername : Deletes an instance of User that matches a given username
func (service *UserService) DeleteByUsername(username string) (models.ServiceResponse, error) {
	if err := service.repo.DeleteUserByUsername(username); err != nil {
		return models.ServiceResponse{Err: fmt.Sprintf("Failed to delete user with username %s", username), Status: 500}, err
	}

	return models.ServiceResponse{Status: 200}, nil
}
