package postgres

import (
	"database/sql"
	"fmt"

	//for connecting to db
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func connect() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected!")
	return db, nil
}
