# Golang Starter Kit

Starter kit lengkap untuk aplikasi web backend menggunakan Go (Golang) dengan fitur-fitur modern dan best practices.

## ✨ Fitur

- 🔐 **Autentikasi JWT** - Login dan logout dengan JSON Web Token
- 👥 **CRUD Users** - Operasi Create, Read, Update, Delete untuk pengguna
- 🗄️ **PostgreSQL** - Database relational yang powerful
- 📦 **Migration** - Database migration untuk versioning schema
- 🏗️ **Clean Architecture** - Pemisahan layer yang jelas (Model, Repository, Service, Controller)
- 🛡️ **Middleware** - Authentication, CORS, dan Logging
- 📝 **Validation** - Input validation menggunakan validator
- 🔒 **Password Hashing** - Bcrypt untuk keamanan password
- 📱 **RESTful API** - API yang mengikuti standar REST

## 🏗️ Struktur Project

```
golang-starter-kit/
├── cmd/
│   └── server/
│       └── main.go              # Entry point aplikasi
├── config/
│   ├── config.go                # Konfigurasi aplikasi
│   └── database.go              # Konfigurasi database
├── database/
│   └── migrations/              # File migration database
│       ├── 001_create_users_table.up.sql
│       └── 001_create_users_table.down.sql
├── internal/
│   ├── controller/              # HTTP handlers
│   │   ├── auth_controller.go
│   │   └── user_controller.go
│   ├── middleware/              # HTTP middleware
│   │   ├── auth.go
│   │   └── cors.go
│   ├── models/                  # Data models
│   │   ├── auth.go
│   │   └── user.go
│   ├── repository/              # Data access layer
│   │   └── user_repository.go
│   ├── routes/                  # Route definitions
│   │   └── routes.go
│   └── service/                 # Business logic layer
│       └── user_service.go
├── pkg/
│   └── utils/                   # Utility functions
│       ├── helpers.go
│       ├── jwt.go
│       └── password.go
├── .env.example                 # Template file environment
├── .gitignore
├── go.mod
└── README.md
```

## 🚀 Cara Menjalankan

### Prerequisites

1. **Go 1.21+** - [Download Go](https://golang.org/dl/)
2. **PostgreSQL** - [Download PostgreSQL](https://www.postgresql.org/download/)
3. **Git** - [Download Git](https://git-scm.com/downloads)

### Setup Database

1. Buat database PostgreSQL:
```sql
CREATE DATABASE golang_starter_kit;
```

2. Buat user database (opsional):
```sql
CREATE USER your_username WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE golang_starter_kit TO your_username;
```

### Setup Project

1. Clone repository:
```bash
git clone <repository-url>
cd golang-starter-kit
```

2. Copy file environment:
```bash
copy .env.example .env
```

3. Edit file `.env` sesuai konfigurasi database Anda:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=golang_starter_kit
DB_SSLMODE=disable

PORT=8080
JWT_SECRET=your_super_secret_jwt_key_here_change_this_in_production
ENV=development
```

4. Install dependencies:
```bash
go mod tidy
```

5. Jalankan migration (manual):
```bash
# Connect ke PostgreSQL dan jalankan file migration
psql -U your_username -d golang_starter_kit -f database/migrations/001_create_users_table.up.sql
```

6. Jalankan aplikasi:
```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:8080`

## 📚 API Documentation

### Health Check
```http
GET /health
```

### Authentication

#### Register
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
}
```

#### Logout
```http
POST /api/v1/auth/logout
Authorization: Bearer <token>
```

### Users Management

#### Get All Users
```http
GET /api/v1/users?page=1&limit=10
```

#### Get User by ID
```http
GET /api/v1/users/1
```

#### Create User
```http
POST /api/v1/users
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "password123"
}
```

#### Update User
```http
PUT /api/v1/users/1
Content-Type: application/json

{
  "name": "Jane Smith",
  "email": "jane.smith@example.com"
}
```

#### Delete User
```http
DELETE /api/v1/users/1
```

### Profile (Protected Routes)

#### Get Profile
```http
GET /api/v1/profile
Authorization: Bearer <token>
```

#### Update Profile
```http
PUT /api/v1/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Updated Name",
  "email": "updated@example.com"
}
```

## 🧪 Testing dengan Postman/curl

### 1. Register User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 2. Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 3. Access Protected Route
```bash
curl -X GET http://localhost:8080/api/v1/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## 🛠️ Dependencies

- **Gin** - HTTP web framework
- **GORM** - ORM library untuk Go
- **PostgreSQL Driver** - Driver database PostgreSQL
- **JWT** - JSON Web Token untuk authentication
- **Bcrypt** - Password hashing
- **Validator** - Input validation
- **Godotenv** - Environment variable loader
- **Migrate** - Database migration tool

## 📁 Layer Architecture

### 1. **Models** (`internal/models/`)
Berisi definisi struct untuk data model dan request/response

### 2. **Repository** (`internal/repository/`)
Layer untuk akses data dan operasi database

### 3. **Service** (`internal/service/`)
Layer untuk business logic dan aturan bisnis

### 4. **Controller** (`internal/controller/`)
Layer untuk handling HTTP request dan response

### 5. **Middleware** (`internal/middleware/`)
Layer untuk middleware seperti authentication, logging, dll

### 6. **Routes** (`internal/routes/`)
Definisi routing dan endpoint API

### 7. **Utils** (`pkg/utils/`)
Utility functions yang bisa digunakan di berbagai layer

## 🔒 Security Features

- **Password Hashing** menggunakan bcrypt
- **JWT Authentication** untuk stateless authentication
- **Input Validation** untuk mencegah invalid data
- **CORS** configuration untuk cross-origin requests
- **Soft Delete** untuk data user

## 🚀 Development Tips

1. **Environment Variables**: Selalu gunakan environment variables untuk konfigurasi sensitif
2. **Error Handling**: Gunakan proper error handling di setiap layer
3. **Validation**: Validasi input di controller layer
4. **Logging**: Gunakan middleware logging untuk monitoring
5. **Database**: Gunakan migration untuk perubahan schema database

## 📝 TODO / Pengembangan Selanjutnya

- [ ] Unit Testing
- [ ] Docker containerization
- [ ] API Rate Limiting
- [ ] Pagination improvement
- [ ] Role-based access control (RBAC)
- [ ] Email verification
- [ ] Password reset functionality
- [ ] API documentation dengan Swagger
- [ ] Monitoring dan metrics
- [ ] Caching layer

## 🤝 Contributing

1. Fork project ini
2. Buat feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

## 👨‍💻 Author

Dibuat dengan ❤️ untuk membantu developer Go pemula memahami clean architecture dan best practices dalam pengembangan API.
