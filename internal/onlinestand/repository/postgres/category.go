package postgres

import (
	"database/sql"

	"github.com/csothen/onlinestand/pkg/models"
)

// CategoryRepository : Postgres Repository that handles transactions related to the Category model
type CategoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository : Creates a new instance of CategoryRepository with an open connection to the Postgres database
func NewCategoryRepository() *CategoryRepository {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	return &CategoryRepository{db}
}

// CreateCategory : Persists a new instance of the model Category in the database
func (repo *CategoryRepository) CreateCategory(category models.Category) (int64, error) {
	panic("Not implemented!")
}

// GetAllCategory : Retrieves all instances of the model Category in the database
func (repo *CategoryRepository) GetAllCategory() ([]*models.Category, error) {
	panic("Not implemented!")
}
