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
func (repo *UserRepository) CreateUser(user models.User) (int, error) {
	var id int

	row := repo.db.QueryRow(`insert into stand_user (email, username, first_name, last_name, password, status, location_id) values(?) returning id;`)
	err := row.Scan(&id)

	return id, err
}

// GetAllUsers : Retrieves all instances of the model User in the database
func (repo *UserRepository) GetAllUsers() ([]*models.User, error) {
	users := make([]*models.User, 0)

	rows, err := repo.db.Query("select stand_user.id, stand_user.email, stand_user.username, stand_user.first_name, stand_user.last_name, stand_user.status, stand_user.location_id, location.country, location.city from stand_user inner join location on stand_user.location_id = location.id")
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID : Retrieves an instance of user that matches a given ID
// or nil in case there is no match
func (repo *UserRepository) GetUserByID(id string) (*models.User, error) {
	row := repo.db.QueryRow("select stand_user.id, stand_user.email, stand_user.username, stand_user.first_name, stand_user.last_name, stand_user.status, stand_user.location_id, location.country, location.city from stand_user inner join location on stand_user.location_id = location.id where stand_user.id=$1", id)

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
	row := repo.db.QueryRow("select stand_user.id, stand_user.email, stand_user.username, stand_user.first_name, stand_user.last_name, stand_user.status, stand_user.location_id, location.country, location.city from stand_user inner join location on stand_user.location_id = location.id where stand_user.username=$1", username)

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
	stmt, err := repo.db.Prepare(`delete from stand_user where id=$1;`)
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
	stmt, err := repo.db.Prepare(`delete from stand_user where username=$1;`)
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
	stmt, err := repo.db.Prepare(`update stand_user set username=?, first_name=?, last_name=? where id=?;`)
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
