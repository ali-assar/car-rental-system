```markdown
# Go Car Rental System

## Overview
GoCarRental is a Go-based application that provides an API for managing car rental operations, including user authentication, agency management, car listing, and reservation handling. This document provides essential information for setting up, using, and contributing to the project.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Dependencies](#dependencies)
- [Makefile](#makefile)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Installation
Ensure you have Go installed on your machine before proceeding.

1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/GoCarRental.git
    cd GoCarRental
    ```

2. Create a MongoDB instance and obtain a connection URL.

3. Create a `.env` file in the project root and set the following variables:
    ```env
    HTTP_LISTEN_ADDRESS=:5000
    JWT_SECRET=your_secret_key_here
    MONGO_DB_NAME=car-rental-system
    MONGO_DB_URL=mongodb://localhost:27017
    MONGO_DB_URL_TEST=mongodb://localhost:27017
    ```

4. Install dependencies and run the application:
    ```bash
    go run main.go
    ```

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

## Project Structure
Briefly describe the high-level project structure and the purpose of each directory or package.

├── api/            # API handlers
├── db/             # Database models and stores
├── middleware/     # Custom middleware (if any)
├── .env            # Environment variable configuration
├── main.go         # Main application entry point
└── ...

## Makefile
The Makefile provides convenient commands for building, running, and testing the Car Rental System. Use the following commands:

```bash
make build     # Build the API binary
make run       # Run the compiled API binary
make seed      # Seed the database with initial data
make docker    # Build and run the API inside a Docker container
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

## Contributing
If you would like to contribute to the project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and ensure they pass any existing tests.
4. Submit a pull request with a clear description of your changes.

---

Feel free to adjust the instructions to better match your project's specifics. Providing clear instructions on configuring the environment, building, running, and contributing to your car rental project is crucial for users and contributors.
```
