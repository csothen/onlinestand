package models

// Feature : Model defining a vehicle's feature
type Feature struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// FeatureService : interface for Feature model that defines the operations on it
type FeatureService interface {
	// Create a feature
	CreateFeature() error

	// Get a single feature based on an id
	GetFeature(id int) (*Feature, error)

	// Get all features
	GetAllFeature() ([]*Feature, error)
}
