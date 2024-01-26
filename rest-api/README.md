```markdown
# REST API Microservice

The REST API microservice serves as the core functionality for user management, authentication, and car reservations in the car rental system.

## Overview

The REST API microservice provides endpoints for creating user accounts, authentication, and reserving cars. It acts as the central component for user interactions within the car rental system.

## Usage
The application exposes RESTful APIs for managing various aspects of the car rental system. Below are some examples of how to interact with the API:

- **User Authentication:**
  - Endpoint: `/api/auth`
  - Method: POST
  - Payload: Provide user credentials to obtain a JSON Web Token (JWT).

- **User Management:**
  - Endpoint: `/api/v1/user`
  - Methods: GET, POST, PUT, DELETE
  - Manage user information.

- **Agency Management:**
  - Endpoint: `/api/v1/agency`
  - Methods: GET
  - Retrieve information about agencies and associated cars.

- **Car Management:**
  - Endpoint: `/api/v1/car`
  - Methods: GET, POST
  - Retrieve a list of cars, reserve a car.

- **Reservation Management:**
  - Endpoint: `/api/v1/reservation`
  - Methods: GET, GET (by ID), GET (cancel by ID)
  - Retrieve reservation information and cancel a reservation.

Refer to the source code for a complete list of endpoints and their functionality.

## Makefile
The Makefile provides convenient commands for building, running, and testing the Car Rental System. Use the following commands:

```bash
make api
make seed      # Seed the database with initial data
make test      # Run tests for the entire project
```

## Environment Variables

GoCarRental uses environment variables for configuration. Before running the application, make sure to set up a `.env` file in the root directory with the following variables:

```env
HTTP_LISTEN_ADDRESS=:5000
JWT_SECRET=your_secret_key_here
MONGO_DB_NAME=car-rental-system
MONGO_DB_URL=mongodb://localhost:27017
MONGO_DB_URL_TEST=mongodb://localhost:27017
```

- **HTTP_LISTEN_ADDRESS**: The address and port on which the API server will listen. Default is `:5000`.

- **JWT_SECRET**: A secret key used for generating and verifying JSON Web Tokens (JWT) for user authentication.

- **MONGO_DB_NAME**: The name of the MongoDB database used by the application.

- **MONGO_DB_URL**: The connection URL for the MongoDB instance used in the development environment.

- **MONGO_DB_URL_TEST**: The connection URL for the MongoDB instance used in the testing environment.

Make sure to adjust these values based on your specific environment and requirements.

## Testing
To test handler functions, check the `api` folder for relevant test files.

