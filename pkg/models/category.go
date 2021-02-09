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
	CreateCategory() error

	// Get a single category based on an id
	GetCategory(id int) (*Category, error)

	// Get all categories
	GetAllCategory() ([]*Category, error)
}
