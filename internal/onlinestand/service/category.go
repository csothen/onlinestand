package service

import (
	"io"

	"github.com/csothen/onlinestand/internal/onlinestand/repository/postgres"
	"github.com/csothen/onlinestand/pkg/models"
)

// CategoryService : Implementation of the interface CategoryService
type CategoryService struct {
	repo models.CategoryRepository
}

// NewCategoryService : Creates a new instance of CategoryService that handles business logic related to the Category model
func NewCategoryService() *CategoryService {
	repo := postgres.NewCategoryRepository()

	return &CategoryService{repo}
}

// GetAll : Retrieves all instances of the model Category
func (service *CategoryService) GetAll() (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// GetByID : Retrieves an instance of the model Category based on an id
func (service *CategoryService) GetByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// Create : Creates and persists a new instance of the model Category
func (service *CategoryService) Create(body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// UpdateByID : Updates and persists an instance of the model Category based on an id
func (service *CategoryService) UpdateByID(id string, body io.ReadCloser) (models.ServiceResponse, error) {
	panic("Not implemented!")
}

// DeleteByID : Deletes an instance of the model Category based on an id
func (service *CategoryService) DeleteByID(id string) (models.ServiceResponse, error) {
	panic("Not implemented!")
}
