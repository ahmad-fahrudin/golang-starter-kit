# Golang Starter Kit

A starter kit for building REST APIs with Golang, Gin, and GORM.

## Features

- REST API with Gin framework
- PostgreSQL database with GORM
- JWT authentication
- User management
- Swagger documentation
- Rate limiting
- CORS support
- Input validation

## Getting Started

### Prerequisites

- Go 1.24+
- PostgreSQL

### Installation

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment variables
3. Run the application:

```bash
go run main.go
```

### Available Commands

- `go run main.go` - Run the server
- `go run main.go server` - Run the server
- `go run main.go migrate` - Run migrations only
- `go run main.go seed` - Run seeders only
- `go run main.go migrate seed` - Run migrations and seeders

## API Documentation

The API documentation is available via Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

### Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

### Endpoints

#### Authentication
- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user
- `POST /api/v1/auth/logout` - Logout user

#### Users
- `POST /api/v1/users` - Create user
- `GET /api/v1/users` - Get all users (paginated)
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

#### Profile (Protected)
- `GET /api/v1/profile` - Get current user profile
- `PUT /api/v1/profile` - Update current user profile

## Project Structure

```
├── config/           # Configuration files
├── database/         # Database migrations and seeders
├── internal/
│   ├── controller/   # HTTP controllers
│   ├── middleware/   # HTTP middlewares
│   ├── models/       # Data models
│   ├── repository/   # Data repositories
│   ├── routes/       # Route definitions
│   └── service/      # Business logic
├── pkg/
│   └── utils/        # Utility functions
├── docs/             # Swagger documentation
└── main.go           # Application entry point
```

## Development

### Adding New Endpoints

1. Define your models in `internal/models/`
2. Create repository methods in `internal/repository/`
3. Implement business logic in `internal/service/`
4. Create controller methods in `internal/controller/`
5. Add routes in `internal/routes/routes.go`
6. Add Swagger annotations to your controller methods
7. Run `swag init` to regenerate documentation

### Updating Swagger Documentation

After adding or modifying API endpoints, regenerate the Swagger documentation:

```bash
swag init
```

## License

This project is licensed under the Apache 2.0 License.
