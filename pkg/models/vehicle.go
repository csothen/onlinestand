package models

// Vehicle : Model defining a vehicle
type Vehicle struct {
	ID                    int      `json:"id"`
	LicensePlate          string   `json:"licensePlate"`
	FirstLicensePlateDate string   `json:"firstLicensePlateDate"`
	Brand                 string   `json:"brand"`
	Model                 string   `json:"model"`
	Kilometers            int      `json:"kilometers"`
	Location              Location `json:"location"`
	Price                 int      `json:"price"`
	Status                string   `json:"status"`
	Negotiable            bool     `json:"negotiable"`
	Condition             string   `json:"condition"`
	Features              []string `json:"features"`
	FuelType              string   `json:"fuelType"`
	Gearbox               string   `json:"gearbox"`
	CubicCapacity         int      `json:"cubicCapacity"`
	EnginePower           int      `json:"enginePower"`
	Seats                 int      `json:"seats"`
	Doors                 int      `json:"doors"`
	Category              string   `json:"category"`
	ExteriorColor         string   `json:"exteriorColor"`
	SeatMaterial          string   `json:"seatMaterial"`
	InteriorColor         string   `json:"interiorColor"`
	ConstructionYear      string   `json:"constructionYear"`
}

// VehicleService : interface for Vehicle model that defines the operations on it
type VehicleService interface {
	// Create a vehicle
	CreateVehicle() error

	// Get a single vehicle based on an id
	GetVehicle(id int) (*Vehicle, error)

	// Get all vehicles
	GetAllVehicle() ([]*Vehicle, error)

	// Get all available vehicles
	GetAllAvailableVehicle() ([]*Vehicle, error)
}
