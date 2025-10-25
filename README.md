# 🎓 Students API

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.24.6-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen?style=for-the-badge)
![Coverage](https://img.shields.io/badge/coverage-0%25-red?style=for-the-badge)

A lightweight, RESTful API for managing student records built with Go and SQLite.

[Features](#-features) • [Quick Start](#-quick-start) • [API Documentation](#-api-documentation) • [Architecture](#-architecture) • [Contributing](#-contributing)

</div>

---

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Quick Start](#-quick-start)
- [Configuration](#-configuration)
- [API Documentation](#-api-documentation)
- [Architecture](#-architecture)
- [Development](#-development)
- [Roadmap](#-roadmap)
- [Contributing](#-contributing)
- [License](#-license)

## ✨ Features

- ✅ **CRUD Operations** - Create, Read, Update, Delete student records
- ✅ **RESTful Design** - Clean and intuitive API endpoints
- ✅ **Input Validation** - Request validation using go-playground/validator
- ✅ **Graceful Shutdown** - Proper signal handling and cleanup
- ✅ **Structured Logging** - JSON-based logging with slog
- ✅ **Configuration Management** - YAML-based config with environment overrides
- ✅ **SQLite Database** - Lightweight, embedded database
- ✅ **Clean Architecture** - Separation of concerns with layered design

## 🛠️ Tech Stack

| Category | Technology |
|----------|-----------|
| **Language** | Go 1.24.6 |
| **Database** | SQLite 3 |
| **Router** | net/http (stdlib) |
| **Validation** | go-playground/validator/v10 |
| **Config** | cleanenv |
| **Logging** | log/slog |

## 📁 Project Structure

```
Students-API/
├── cmd/
│   └── students-api/
│       └── main.go              # Application entry point
├── config/
│   └── local.yaml               # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go            # Config loader
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go   # HTTP handlers
│   ├── storage/
│   │   ├── storage.go           # Storage interface
│   │   └── sqlite/
│   │       └── sqlite.go        # SQLite implementation
│   ├── types/
│   │   └── types.go             # Domain models
│   └── utils/
│       └── response/
│           └── response.go      # Response helpers
├── storage/
│   └── storage.db               # SQLite database file
├── go.mod
├── go.sum
├── README.md
└── IMPROVEMENTS.md              # Future enhancements
```

### 🏗️ Architecture Layers

```
┌─────────────────────────────────────┐
│         HTTP Handlers               │  ← Request/Response handling
├─────────────────────────────────────┤
│         Storage Interface           │  ← Abstraction layer
├─────────────────────────────────────┤
│      SQLite Implementation          │  ← Database operations
└─────────────────────────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- Go 1.24.6 or higher
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Flack74/Students-API.git
   cd Students-API
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Create configuration file**
   ```bash
   # config/local.yaml already exists with default settings
   cat config/local.yaml
   ```

4. **Run the application**
   ```bash
   # Using config file path
   go run cmd/students-api/main.go -config=config/local.yaml
   
   # Or using environment variable
   export Config_Path=config/local.yaml
   go run cmd/students-api/main.go
   ```

5. **Verify it's running**
   ```bash
   curl http://localhost:8082/api/students
   ```

### 🐳 Docker (Coming Soon)

```bash
docker build -t students-api .
docker run -p 8082:8082 students-api
```

## ⚙️ Configuration

Configuration is managed via YAML files. Create a config file with the following structure:

```yaml
# config/local.yaml
env: "dev"                          # Environment: dev, staging, prod
storage_path: "storage/storage.db"  # SQLite database path
http_server:
  address: "localhost:8082"         # Server address
```

### Environment Variables

You can override config file location using:

```bash
export Config_Path=/path/to/config.yaml
```

## 📚 API Documentation

### Base URL
```
http://localhost:8082/api
```

### Endpoints

#### Create Student
```http
POST /api/students
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 20
}
```

**Response (201 Created)**
```json
{
  "id": 1
}
```

#### Get Student by ID
```http
GET /api/students/{id}
```

**Response (200 OK)**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "age": 20
}
```

#### Get All Students
```http
GET /api/students
```

**Response (200 OK)**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 20
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane@example.com",
    "age": 22
  }
]
```

#### Update Student
```http
PUT /api/students/{id}
Content-Type: application/json

{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "age": 21
}
```

**Response (201 Created)**
```json
{
  "message": "student upated successfully"
}
```

#### Delete Student
```http
DELETE /api/students/{id}
```

**Response (200 OK)**
```json
{
  "message": "student deleted successfully"
}
```

### Error Responses

#### Validation Error (400 Bad Request)
```json
{
  "status": "Error",
  "error": "field Name is required field, field Email is required field"
}
```

#### General Error (500 Internal Server Error)
```json
{
  "status": "Error",
  "error": "database connection failed"
}
```

### Request Validation Rules

| Field | Type | Required | Validation |
|-------|------|----------|------------|
| name | string | ✅ | Must not be empty |
| email | string | ✅ | Must not be empty |
| age | integer | ✅ | Must be provided |

## 🏛️ Architecture

### Design Principles

1. **Clean Architecture** - Separation of concerns with clear boundaries
2. **Dependency Injection** - Storage interface injected into handlers
3. **Interface-Based Design** - Easy to swap implementations
4. **Single Responsibility** - Each package has one clear purpose

### Data Flow

```
HTTP Request
    ↓
Handler (Validation)
    ↓
Storage Interface
    ↓
SQLite Implementation
    ↓
Database
    ↓
Response
```

### Key Components

#### 1. Handlers (`internal/http/handlers/student/`)
- HTTP request/response handling
- Input validation
- Error handling
- Response formatting

#### 2. Storage Interface (`internal/storage/`)
- Defines contract for data operations
- Allows multiple implementations (SQLite, PostgreSQL, etc.)

#### 3. SQLite Implementation (`internal/storage/sqlite/`)
- Concrete implementation of Storage interface
- Database connection management
- SQL query execution

#### 4. Types (`internal/types/`)
- Domain models
- Data structures

#### 5. Utils (`internal/utils/`)
- Response helpers
- Common utilities

## 💻 Development

### Building

```bash
# Build binary
go build -o bin/students-api cmd/students-api/main.go

# Run binary
./bin/students-api -config=config/local.yaml
```

### Testing

```bash
# Run tests (coming soon)
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...
```

### Linting

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

### Database Management

The SQLite database is automatically created on first run. To reset:

```bash
rm storage/storage.db
# Restart the application
```

### Example Requests

```bash
# Create a student
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com","age":19}'

# Get all students
curl http://localhost:8082/api/students

# Get student by ID
curl http://localhost:8082/api/students/1

# Update student
curl -X PUT http://localhost:8082/api/students/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice.new@example.com","age":20}'

# Delete student
curl -X DELETE http://localhost:8082/api/students/1
```

## 🗺️ Roadmap

See [IMPROVEMENTS.md](IMPROVEMENTS.md) for detailed enhancement plans.

### Phase 1 - Security & Stability
- [ ] Fix XSS vulnerabilities
- [ ] Improve error handling
- [ ] Add database connection pooling
- [ ] Input sanitization

### Phase 2 - Features
- [ ] Pagination and filtering
- [ ] Search functionality
- [ ] PATCH endpoint for partial updates
- [ ] Bulk operations

### Phase 3 - Production Ready
- [ ] Authentication & authorization
- [ ] Rate limiting
- [ ] Caching layer
- [ ] Metrics and monitoring
- [ ] Docker support
- [ ] CI/CD pipeline

### Phase 4 - Advanced
- [ ] GraphQL API
- [ ] WebSocket support
- [ ] Multi-tenancy
- [ ] Advanced analytics

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and idioms
- Write tests for new features
- Update documentation
- Run linters before committing
- Keep commits atomic and descriptive

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**Flack74**
- GitHub: [@Flack74](https://github.com/Flack74)

## 🙏 Acknowledgments

- [cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Configuration management
- [validator](https://github.com/go-playground/validator) - Input validation
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver

## 📞 Support

If you have any questions or need help, please:
- Open an issue on GitHub
- Check existing issues for solutions
- Review the [IMPROVEMENTS.md](IMPROVEMENTS.md) for known limitations

---

<div align="center">

**⭐ Star this repo if you find it helpful!**

Made with ❤️ using Go

</div>
