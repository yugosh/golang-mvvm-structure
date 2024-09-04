# Backend Golang MVVM Architecture with (RabbitMQ, Redis, SocketIO, NodeJS/ExpressJS)

This project is designed to implement a scalable backend architecture using the MVVM (Model-View-ViewModel) pattern in Golang. The architecture is built to enhance modularity, maintainability, and scalability for modern backend services. Key components include integration with RabbitMQ for message queuing, Redis for in-memory data storage, and Socket.IO for real-time communication. NodeJS (with ExpressJS) is also used as an additional layer for specific services such as handling Socket.IO and serving real-time data.

## Project Structure

```bash
/BACKENG-GOLANG-MVVM
├── DockerFile                     # Dockerfile for the Go application (Golang backend)
├── Makefile                       # Makefile for automation tasks like build and other commands
├── README.md                      # Project documentation
├── attachment                     # Folder to store attachment files, such as images or PDFs
│   ├── images01.jpg               # Example image in the attachment folder
│   └── test.pdf                   # Example PDF file in the attachment folder
├── cmd                            # Main folder for the Go application (entry point of the app)
│   └── app
│       └── main.go                # Main Go file to run the application
├── docker-compose.yml             # Docker Compose file to orchestrate multiple services (Go, Node.js, Redis, RabbitMQ, Azure SQL)
├── go.mod                         # Go module file for managing dependencies
├── go.sum                         # Go checksum file to ensure dependency integrity
├── internal                       # Internal folder for application logic
│   ├── app                        # Application folder containing controllers, services, and models
│   │   ├── controllers            # Folder for controller logic
│   │   │   ├── auth_controller.go # Authentication-related controller
│   │   │   ├── base_controller.go # Base controller for common functionalities
│   │   │   ├── notification_controller.go # Controller for handling notifications
│   │   │   └── user_controller.go  # Controller for user-related functionalities
│   │   ├── mappers                 # Folder for data mapping logic
│   │   │   └── user_mapper.go      # Mapper for user-related data transformations
│   │   ├── middleware              # Folder for middlewares
│   │   │   └── cors_middleware.go  # Middleware for handling CORS
│   │   ├── models                  # Folder for defining models
│   │   │   └── user.go             # User model definition
│   │   ├── repositories            # Folder for repository logic
│   │   │   └── user_repository.go  # User repository for data persistence
│   │   ├── routes                  # Folder for route definitions
│   │   │   └── routes.go           # Main route file
│   │   ├── services                # Folder for business logic and services
│   │   │   ├── consumer.go         # RabbitMQ consumer service
│   │   │   ├── firebase.go         # Service for Firebase interactions
│   │   │   ├── producer.go         # RabbitMQ producer service
│   │   │   ├── socketio.go         # Service for handling Socket.IO logic
│   │   └── viewmodels              # Folder for view models
│       │   └── user_viewmodel.go   # View model for user-related data
│   ├── config                      # Configuration files for the application
│   │   ├── rabbitmq.go             # RabbitMQ configuration
│   │   ├── redis.go                # Redis configuration
│   │   └── serviceAccountKey.json  # Firebase service account key for push notifications
│   └── db                          # Database-related logic and migrations
│       ├── db.go                   # Main database initialization file
│       └── migrations              # Folder for database migrations
├── pkg                            # Folder for utilities and reusable components
│   └── utils                       # Utility functions
├── socketio                        # Folder for the Socket.IO Node.js service
│   ├── Dockerfile                  # Dockerfile for the Node.js application
│   ├── index.js                    # Main entry point for the Socket.IO server
│   └── package.json                # Package.json for managing Node.js dependencies
├── static                          # Folder for static HTML and frontend files
│   ├── dashboard.html              # Dashboard HTML page
│   ├── firebase-messaging-sw.js    # Firebase service worker file for notifications
│   ├── formula.html                # Formula page HTML file
│   ├── index.html                  # Main landing page for the application
│   ├── javascript.js               # JavaScript file for handling client-side logic
│   ├── notification.html           # Notification page HTML file
│   └── register.html               # Registration page HTML file
```

## Getting Started

## RUN USING LOCAL

Clone the repository:

```bash
git clone https://github.com/yugosh/backend-golang-mvvm.git
```

Install node_modules nodejs
```bash
cd socketio
```

```bash
npm i
```

Install dependencies:
```bash
cd backend-golang-mvvm
```

```bash
go mod download
```

Change .ENV variable APP_ENV to 'blank'
```bash
APP_ENV=local
```

Run database migrations:
``` bash
go run internal/db/migrations.go
```

Run the application:
```bash
go run cmd/app/main.go
```

## RUN USING DOCKER

Change .ENV variable APP_ENV to 'docker'
```bash
APP_ENV=docker
```

Build Docker:
```bash
docker-compose up --build
```

# Contributing
Contributions are welcome! Please fork this repository and submit a pull request with your improvements.

# License
This project is licensed under the MIT License - see the LICENSE file for details.
