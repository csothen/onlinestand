#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    CREATE TABLE IF NOT EXISTS location(
      id SERIAL PRIMARY KEY,
      country VARCHAR(20) NOT NULL,
      city VARCHAR(20) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS category (
      id SERIAL PRIMARY KEY,
      code VARCHAR(10) NOT NULL,
      short_description VARCHAR(40) NOT NULL,
      long_description TEXT NOT NULL
    );

    CREATE TABLE IF NOT EXISTS feature (
        id SERIAL PRIMARY KEY,
        name VARCHAR(20),
        description TEXT
    );

    CREATE TABLE IF NOT EXISTS stand_user (
	  id SERIAL PRIMARY KEY,
      email VARCHAR(50) NOT NULL,
      username VARCHAR(20) NOT NULL,
      password VARCHAR(60) NOT NULL,
	  first_name VARCHAR(20) NOT NULL,
      last_name VARCHAR(20) NOT NULL,
	  status VARCHAR(20) NOT NULL,
	  location_id INT references location(id) ,
	  UNIQUE(email, username)
	);

    CREATE TABLE IF NOT EXISTS vehicle (
        id SERIAL PRIMARY KEY,
        license_plate VARCHAR(20) NOT NULL,
        first_license_plate_date TIMESTAMP,
        brand VARCHAR(20),
        model VARCHAR(20),
        kilometers INT NOT NULL,
        price MONEY NOT NULL,
        status VARCHAR(20) NOT NULL,
        negotiable BOOLEAN NOT NULL,
        condition VARCHAR(20),
        fuel_type VARCHAR(20),
        gearbox VARCHAR(20),
        cubic_capacity INT,
        engine_power INT,
        seats INT,
        doors INT,
        exterior_color VARCHAR(20),
        seat_material VARCHAR(20),
        interior_color VARCHAR(20),
        construction_year INT,
        category_id INT references category(id),
        location_id INT references location(id),
        UNIQUE(license_plate)
    );

    CREATE TABLE IF NOT EXISTS vehicle_feature (
        vehicle_id INT references vehicle(id),
        feature_id INT references feature(id),
        PRIMARY KEY(vehicle_id, feature_id)
    );

	CREATE INDEX idx_stand_user_location_id ON stand_user (location_id);
    CREATE INDEX idx_vehicle_category_id ON vehicle (category_id);
  COMMIT;
EOSQL