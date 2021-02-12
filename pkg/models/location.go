package models

// Location : Model defining a location in a certain country
type Location struct {
	ID      int    `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
}

// LocationService : interface for location defining operations on it
type LocationService interface {
	// Create a location
	CreateLocation(l Location) ServiceResponse

	// Get all locations
	GetAllLocation() ServiceResponse
}

// LocationRepository : interface for Location model that defines operations available in the repository
type LocationRepository interface {
	// Create a location
	CreateLocation(l Location) error

	// Get all locations
	GetAllLocation() ([]*Location, error)
}
