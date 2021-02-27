package postgres

import (
	"database/sql"

	"github.com/csothen/onlinestand/pkg/models"
)

// UserRepository : Postgres Repository that handles transactions related to the User model
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository : Creates a new instance of UserRepository with an open connection to the Postgres database
func NewUserRepository() *UserRepository {
	db, err := connect()
	if err != nil {
		panic(err)
	}

	return &UserRepository{db}
}

// CreateUser : Persists a new instance of the model User in the database
func (repo *UserRepository) CreateUser(user models.User) (int64, error) {
	stmt, err := repo.db.Prepare(`insert into user (email, username, first_name, last_name, password, status, location_id) values(?);`)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Email, user.Username, user.FirstName, user.LastName, user.Password, user.Status, user.Location.ID)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

// GetAllUsers : Retrieves all instances of the model User in the database
func (repo *UserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	rows, err := repo.db.Query("select user.id, user.email, user.username, user.first_name, user.last_name, user.status, user.location_id, location.country, location.city from user inner join location on user.location_id = location.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(user.ID, user.Email, user.Username, user.FirstName, user.LastName, user.Status, user.Location.ID, user.Location.Country, user.Location.City)
		users = append(users, user)
		if err != nil {
			return nil, err
		}
	}

	if rows.Err() != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID : Retrieves an instance of user that matches a given ID
// or nil in case there is no match
func (repo *UserRepository) GetUserByID(id string) (*models.User, error) {
	row := repo.db.QueryRow("select user.id, user.email, user.username, user.first_name, user.last_name, user.status, user.location_id, location.country, location.city from user inner join location on user.location_id = location.id where user.id=$1", id)

	user := &models.User{}
	err := row.Scan(user.ID, user.Email, user.Username, user.FirstName, user.LastName, user.Status, user.Location.ID, user.Location.Country, user.Location.City)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// GetUserByUsername : Retrieves an instance of user that matches a given username
// or nil in case there is no match
func (repo *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	row := repo.db.QueryRow("select user.id, user.email, user.username, user.first_name, user.last_name, user.status, user.location_id, location.country, location.city from user inner join location on user.location_id = location.id where user.username=$1", username)

	user := &models.User{}
	err := row.Scan(user.ID, user.Email, user.Username, user.FirstName, user.LastName, user.Status, user.Location.ID, user.Location.Country, user.Location.City)

	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return user, nil
	default:
		return nil, err
	}
}

// DeleteUserByID : Deletes an instance of user that matches a given ID
func (repo *UserRepository) DeleteUserByID(id string) error {
	stmt, err := repo.db.Prepare(`delete from user where id=$1;`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserByUsername : Deletes an instance of user that matches a given username
func (repo *UserRepository) DeleteUserByUsername(username string) error {
	stmt, err := repo.db.Prepare(`delete from user where username=$1;`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}

	return nil
}

// Save : Persists an existing instance of the model User in the database
func (repo *UserRepository) Save(user *models.User) error {
	stmt, err := repo.db.Prepare(`update user set username=?, first_name=?, last_name=? where id=?;`)
	defer stmt.Close()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Username, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return err
	}

	return nil
}
