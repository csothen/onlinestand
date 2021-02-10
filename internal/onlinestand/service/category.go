package service

import (
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

// CreateCategory : Creates and persists a new instance of the model Category
func (service *CategoryService) CreateCategory(c models.Category) error {
	panic("Not implemented!")
}

// GetAllCategory : Retrieves all instances of the model Category
func (service *CategoryService) GetAllCategory() ([]*models.Category, error) {
	panic("Not implemented!")
}
