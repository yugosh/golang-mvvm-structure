# Backend Golang MVVM Architecture

This project implements a scalable backend architecture using the MVVM (Model-View-ViewModel) pattern in Golang. The architecture is designed to separate concerns and improve maintainability and scalability.

## Project Structure

```bash
/backend-golang-mvvm
├── cmd
│   └── app
│       └── main.go          # Entry point of the application
├── internal
│   ├── app
│   │   ├── routes           # API routes and middleware grouping
│   │   │   └── routes.go
│   │   ├── middleware       # Middleware for request & response handling
│   │   │   ├── auth.go
│   │   │   ├── logging.go
│   │   │   └── recovery.go
│   │   ├── controllers      # Controllers/Handlers for HTTP requests
│   │   │   └── user_controller.go
│   │   ├── viewmodels       # ViewModels for presentation logic
│   │   │   └── user_viewmodel.go
│   │   ├── models           # Models for data structures and business logic
│   │   │   └── user.go
│   │   ├── services         # Business logic and ORM interactions
│   │   │   └── user_service.go
│   │   ├── repositories     # Data access and CRUD operations
│   │   │   └── user_repository.go
│   │   └── mappers          # Mapping between Models and ViewModels
│   │       └── user_mapper.go
│   ├── config               # Application configuration
│   │   └── config.go
│   └── db
│       ├── migrations       # Database migration scripts
│       │   └── 001_create_users_table.up.sql
│       └── db.go            # Database connection setup
├── pkg                      # Reusable utility packages
│   └── utils
│       ├── logger.go
│       └── json.go
└── go.mod                   # Dependency management
└── go.sum                   # Dependency lock file
└── Makefile                 # Build and run commands
└── Dockerfile               # Docker containerization instructions
└── .env                     # Environment variables
└── README.md                # Project documentation
```

# Explanation of Structure

- cmd/app/main.go:
    - Entry point of the application that initializes the server and sets up routes.

- internal/app/routes:
    - routes.go: Contains all API routes and groups middleware. This file manages which endpoints are directed to which handlers.
- internal/app/middleware:
    - auth.go: Middleware for request authentication.
    - logging.go: Middleware for logging requests and responses.
    - recovery.go: Middleware for recovering the application from panic and returning appropriate error messages.

- internal/app/controllers:
    - user_controller.go: Contains handlers for managing requests related to User. This is the View part of MVVM.

- internal/app/viewmodels:
    - user_viewmodel.go: Manages the presentation logic for User, handling data sent from or to the Controller.

- internal/app/models:
    - user.go: Defines the User entity, including properties and possibly basic validations. This is the Model part of MVVM.

- internal/app/services:
    - user_service.go: Contains business logic and interactions with ORM (e.g., GORM) for User. It serves as the connection between the Model and ViewModel.

- internal/app/repositories:
    - user_repository.go: Handles data access (CRUD operations) with the database. This layer abstracts the ORM, allowing database technology changes without affecting business logic.

- internal/app/mappers:
    - user_mapper.go: Contains logic to map data from Model to ViewModel or vice versa. This is useful for converting raw data from the database into a format ready for presentation in the View.

- internal/config:
    - config.go: Contains application configuration such as database connections, server settings, etc., usually sourced from .env.

- internal/db:
    - migrations/: This directory contains SQL files for database migrations, ensuring the database schema is up-to-date with application requirements.
    - db.go: Manages the database connection and provides setup functions.

- pkg/utils:
    - logger.go: Utility functions for logging that can be used throughout the application.
    - json.go: Utility functions for JSON serialization and deserialization.

- go.mod & go.sum:
    - Manage project dependencies in Golang.

- Dockerfile:
    - Instructions for containerizing the application using Docker.

- .env:
    - Contains environment variables such as database connections, API keys, etc.

- README.md:
    - Project documentation that provides basic information on how to run and develop the application.

# Implementation Steps
- Routes Setup:
    -Define your routes in internal/app/routes/routes.go.
    - Add appropriate middleware to routes that require authentication, logging, etc.

- Middleware Implementation:
    - Implement authentication, logging, and recovery logic in internal/app/middleware.

- Controller/Handler:
    - Create HTTP handlers in internal/app/controllers that interact with the ViewModel.

- ViewModel Logic:
    - Implement the presentation logic in internal/app/viewmodels.

- Model Definition:
    - Define your entities and business logic in internal/app/models.

- Service Layer:
    - Implement business logic and interactions with ORM in internal/app/services.

- Repository Layer:
    - Create abstractions for data access in internal/app/repositories.

- Mapping Model:
    - Implement mapping between Model and ViewModel in internal/app/mappers.


# Getting Started
Clone the repository:

```bash
git clone https://github.com/yourusername/backend-golang-mvvm.git
cd backend-golang-mvvm
```

# Install dependencies:
```bash
go mod download
```

# Run database migrations:
``` bash
go run internal/db/migrations.go
```

# Run the application:
```bash
go run cmd/app/main.go
```

# Build Docker image:
```bash
docker build -t backend-golang-mvvm .
```

# Run the Docker container:
```bash
docker run -p 8080:8080 backend-golang-mvvm
```

# Contributing
Contributions are welcome! Please fork this repository and submit a pull request with your improvements.

# License
This project is licensed under the MIT License - see the LICENSE file for details.