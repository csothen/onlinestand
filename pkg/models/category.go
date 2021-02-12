package models

// Category : Model defining a vehicle's category
type Category struct {
	ID               int    `json:"id"`
	Code             string `json:"code"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
}

// CategoryService : interface for Category model that defines the operations on it
type CategoryService interface {
	// Create a category
	CreateCategory(c Category) ServiceResponse

	// Get all categories
	GetAllCategory() ServiceResponse
}

// CategoryRepository : interface for Category model that defines operations available in the repository
type CategoryRepository interface {
	// Create a category
	CreateCategory(c Category) error

	// Get all categories
	GetAllCategory() ([]*Category, error)
}
