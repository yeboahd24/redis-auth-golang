# Redis Authentication Service

A Go-based authentication service that uses Redis for user management and JWT for session handling.

## Features

- User registration with password strength validation
- Secure login with JWT token generation
- Redis-based user storage
- Password hashing using bcrypt
- Rate limiting middleware
- Configuration management using Viper

## Prerequisites

- Go 1.23 or higher
- Redis server running locally or accessible
- Git (for cloning the repository)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yeboahd24/redis-auth-golang
cd redis-auth
```

2. Install dependencies:
```bash
go mod download
```

3. Configure the application:
   - Copy `config.yaml.example` to `config.yaml`
   - Update the configuration values as needed:
     - `redis_addr`: Redis server address (default: "localhost:6379")
     - `redis_password`: Redis password (if required)
     - `redis_db`: Redis database number
     - `jwt_secret`: Secret key for JWT token generation

## Running the Application

1. Start the Redis server
2. Run the application:
```bash
go run main.go
```

The server will start on port 8080.

## API Endpoints

### POST /signup
Register a new user.

Request body:
```json
{
    "email": "user@example.com",
    "password": "SecurePass123!"
}
```

Password requirements:
- Minimum 8 characters
- At least one uppercase letter
- At least one lowercase letter
- At least one number
- At least one special character

### POST /login
Authenticate a user and receive a JWT token.

Request body:
```json
{
    "email": "user@example.com",
    "password": "SecurePass123!"
}
```

Response:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

## Security Features

- Password hashing using bcrypt
- JWT-based authentication
- Rate limiting (60 requests per minute per IP)
- Input validation
- Secure password requirements

## Project Structure

```
redis-auth/
├── config/
│   └── config.go         # Configuration management
├── handlers/
│   └── auth.go          # HTTP request handlers
├── middleware/
│   └── ratelimit.go     # Rate limiting middleware
├── models/
│   └── user.go          # User data model
├── repository/
│   └── redis.go         # Redis operations
├── utils/
│   ├── errors.go        # Error handling utilities
│   ├── jwt.go           # JWT operations
│   └── password.go      # Password validation
├── config.yaml          # Configuration file
├── go.mod              # Go module file
├── go.sum              # Go module checksums
├── main.go             # Application entry point
└── README.md           # This file
```

## Error Handling

The service returns appropriate HTTP status codes:
- 200: Success
- 201: Resource created
- 400: Bad request (invalid input)
- 401: Unauthorized
- 429: Too many requests
- 500: Internal server error

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

[Add your license information here]