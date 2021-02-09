package models

// User : Model defining a user
type User struct {
	ID        int      `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Location  Location `json:"location"`
	Status    string   `json:"status"`
}

// UserService : interface for User model that defines the operations on it
type UserService interface {
	// Create a user
	CreateUser() error

	// Get a single user based on an id
	GetUser(id int) (*User, error)

	// Get all users
	GetAllUser() (*[]User, error)
}
