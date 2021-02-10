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
	CreateFeature(f Feature) error

	// Get all features
	GetAllFeature() ([]*Feature, error)
}

// FeatureRepository : interface for Feature model that defines operations available in the repository
type FeatureRepository interface {
	// Create a feature
	CreateFeature(f Feature) error

	// Get all categories
	GetAllFeature() ([]*Feature, error)
}
