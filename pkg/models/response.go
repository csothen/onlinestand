package models

// ServiceResponse : Is a structure that holds all the information needed for the controller to know how to behave
type ServiceResponse struct {
	Value  interface{} `json:"value,omitempty"`
	Err    string      `json:"error,omitempty"`
	Status int         `json:"status"`
}
