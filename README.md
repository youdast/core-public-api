# Core Public API

This is a Clean Architecture REST API service built with Go (Golang) and Fiber.

## Tech Stack

- **Language**: Go 1.20+
- **Framework**: [Fiber v2](https://gofiber.io/)
- **Database**: PostgreSQL
- **ORM**: [GORM](https://gorm.io/)
- **Configuration**: [Viper](https://github.com/spf13/viper)

## Project Structure

The project follows the Standard Go Project Layout and Clean Architecture principles:

```text
.
├── cmd/api/         # Main entry point
├── config/          # Configuration loader
├── internal/        # Private application code
│   ├── delivery/    # HTTP Handlers (Fiber)
│   ├── domain/      # Entities and Interfaces
│   ├── repository/  # Data Access Layer (GORM)
│   └── usecase/     # Business Logic
├── pkg/             # Public library code (Database, Utils)
└── .env             # Environment variables
```

## Setup & Installation

1.  **Clone the repository**
2.  **Setup Database**
    - Ensure PostgreSQL is running.
    - Create a database (e.g., `postgres`).
3.  **Configure Environment**
    - Copy `.env.example` to `.env`:
      ```bash
      cp .env.example .env
      ```
    - Update `.env` with your database credentials and SMTP settings.

## Running the Application

### Option 1: Standard Go Run
If Go is in your system PATH:
```bash
go run cmd/api/main.go
```

### Option 2: Using Helper Script (Windows)
If you are using FlyEnv or have PATH issues, use the provided PowerShell script:
```powershell
.\dev.ps1
```

## API Endpoints

### Health Check
- `GET /health` - Check service status

### Users
- `GET /api/v1/users` - List all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create a new user
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com"
  }
  ```
