package models

// Vehicle holds all of a vehicle's information
type Vehicle struct {
	VIN                   string
	LicensePlate          string
	FirstLicensePlateDate string
	Brand                 string
	Model                 string
	Kilometers            int
	Price                 int
	Negotiable            bool
	Sold                  bool
	Condition             string
	Extras                []string
	FuelType              string
	Gearbox               string
	CubicCapacity         int
	EnginePower           int
	EngineEmissions       int
	Seats                 int
	Doors                 int
	VehicleType           string
	VehicleColor          string
	SeatMaterial          string
	InteriorColor         string
	WheelSize             int
	WheelType             string
	ConstructionYear      string
	CreatedOn             string
	UpdatedOn             string
	DeletedOn             string
}
