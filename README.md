# User Management API

A RESTful API for basic user management operations built with Go and MongoDB.

## Features

- Create new users
- Retrieve user details by ID
- Delete users by ID

## Tech Stack

- **Go**: Backend language
- **MongoDB**: Database
- **mongo-go-driver**: Official MongoDB driver for Go

## Prerequisites

- Go 1.16+
- MongoDB 4.4+
- Git

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/sharukh010/user-management.git
   ```

2. Navigate to the project directory:
   ```
   cd user-management
   ```

3. Install dependencies:
   ```
   go mod download
   ```

4. Make sure MongoDB is running locally.

5. Run the application:
   ```
   go run main.go
   ```

## API Endpoints

| Method | Endpoint       | Description       | Request Body | Response        |
|--------|----------------|-------------------|--------------|-----------------|
| POST   | /user          | Create a new user | User object  | Created user    |
| GET    | /user/{userId} | Get user by ID    | None         | User details    |
| DELETE | /user/{userId} | Delete a user     | None         | Status message  |

## Example Requests

### Create User
```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "password": "securepassword"}'
```

### Get User by ID
```bash
curl -X GET http://localhost:8080/user/60d21b4667d0d8992e610c85
```

### Delete User
```bash
curl -X DELETE http://localhost:8080/user/60d21b4667d0d8992e610c85
```

## Project Structure

```
sharukh010-user-management/
├── go.mod           # Go modules
├── go.sum           # Go modules checksum
├── cmd/             # Application entry points
│   └── main/        # Main application
│       └── main.go  # Main entry point
└── pkg/             # Package code
    ├── config/      # Configuration settings
    │   └── config.go
    ├── controllers/ # Request handlers
    │   └── user-controllers.go
    ├── models/      # Data models
    │   └── user.go
    └── routes/      # API routes
        └── user-router.go
```

## Configuration

The application uses environment variables for configuration:
- `MONGODB_URI`: MongoDB connection string (default: "mongodb://localhost:27017")
- `DB_NAME`: Database name (default: "user_management")
- `PORT`: Server port (default: "8001")

## Development

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

Sharukh - [GitHub](https://github.com/sharukh010)

Project Link: [https://github.com/sharukh010/user-management](https://github.com/sharukh010/user-management)