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
	CreateUser(u User) ServiceResponse

	// Get a single user based on an id
	GetUser(id int) ServiceResponse

	// Get all users
	GetAllUser() ServiceResponse
}

// UserRepository : interface for User model that defines operations available in the repository
type UserRepository interface {
	// Create a user
	CreateUser(user User) (*User, error)

	// Get all users
	GetAllUsers() ([]*User, error)

	// Get user by ID
	GetUserByID(id string) (*User, error)

	// Get user by username
	GetUserByUsername(username string) (*User, error)

	// Delete user by ID
	DeleteUserByID(id string) error

	// Delete user by username
	DeleteUserByUsername(username string) error

	// Saves a user
	Save(user *User) (*User, error)
}
