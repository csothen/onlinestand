# Online Stand Web Application

Web Application to manage an online stand based in a microservice architecture

## Models

### User Model

- First Name
- Last Name
- Email
- Username
- Password
- Location (Location of the user)
- Status (If the user is deleted/active/deactivated)

### Vehicle Model

- Identifier
- License Plate
- First License Plate Date
- Brand
- Model
- Kilometers
- Price
- Location (Location where the vehicle is currently holding country and city)
- Negotiable (Boolean if the price is negotiable)
- Status (If the vehicles is available/sold/reserved)
- Condition (If the vehicle's condition is new/renewed/used/broken)
- Features (A series of features that the user can select)
- Fuel Type (Electric/Diesel/Gasoline/Hybrid)
- Gearbox (Automatic/Manual)
- Cubic Capacity
- Engine Power
- Seats
- Doors
- Vehicle Category (From the [available categories](https://en.wikipedia.org/wiki/Vehicle_category))
- Vehicle Color
- Seat Material
- Interior Color
- Construction Year

### Vehicle Category Model

- Identifier
- Code
- Short Description
- Long Description

### Vehicle Feature Model

- Identifier
- Name
- Description

### Location Model

- Identifier
- Country
- City

## Data validations

### User Model Validations

- Email - Must follow the [email regex](https://emailregex.com/)
- Username - Must be between 4 and 16 characters
- Password - Must be atleast 8 characters long, have 1 upper case, 1 lower case and 1 number

### Vehicle Model Validations

- License Plate - The validation must be done depending on the vehicles location
- First License Plate Date - A date larger than the vehicle's construction year and smaller than the current date
- Kilometers - An integer number equal or greater than 0
- Price - A float number equal or greater than 0 in euros
- Location - Must be a valid country and the city must belong to the country selected
- Cubic Capacity - An integer number equal or greater than 0
- Engine Power - An integer number equal or greater than 0
- Seats - An integer number greater than 0
- Doors - An integer number equal or greater than 0
- Construction Year - A date smaller than the current date

### Location Model Validations

- City - Must be a city that exists in the location's country

## Flows

### Login

- Send credentials
- Retrieve user matching username
- Match passwords
- Generate authentication token
- Return authentication token

### Registration

- Send user information
- Run validation checks on the data provided
- Create a new user in the application
- Return the username of the user created

### Consult all available vehicles

- Retrieve all available vehicles with filters
- Query the data with the given filters
- Return paginated vehicle information

### Add a new vehicle

- Insert all information required for a vehicle
- Run validation checks on the data provided
- Create a new vehicle in the application
- Return the identifier of the vehicle created

### Consult a vehicle's information

- Retrieve a vehicle by its identifier
- Query the data with the given identifier
- Return the public information related to the vehicle
