# Note Taking App

A simple API-based note taking application written in Go. The application uses JWT for authentication and follows the REST architecture. 

## Prerequisites

- Golang 1.14 or newer
- PostgreSQL

## Technologies Used

- **Gin**: A HTTP web framework written in Go for routing.
- **JWT**: Used for user authentication.
- **PostgreSQL**: The database used for storing user and note data.

Other dependencies can be found in the `go.mod` file.

## Project Structure

The project follows a layered architecture inspired by Go-Kit. Here's a high-level overview of the structure and responsibility of each directory:

- **cmd**: This is the main entry point of the application. It contains the main function which runs the server.

- **repository**: This directory contains data access code and database related logic. It's where we define our models and database interactions. It also includes database migration scripts.

- **service**: The business logic of the application resides here. This layer interacts with the repository layer to fetch, process, and write data.

- **transport**: This layer handles the HTTP routing and request/response processing. It contains endpoint definitions, middleware, and error handling.

Below is the directory structure for more details:
```
.
├── cmd
│   └── app
│       └── main.go
├── go.mod
├── go.sum
├── README.md
├── repository
│   ├── models.go
│   ├── note_repository.go
│   └── user_repository.go
├── service
│   ├── note_service.go
│   └── user_service.go
└── transport
    ├── note_handler.go
    ├── routing.go
    └── user_handler.go
```

## Setup

1. Clone this repository.
2. Set your database credentials in the `.env` file.
3. Run the command `go run .` in the root folder of the project to start the server.

## Endpoints

- **Signup** - `POST /signup`
    - Request Body: `{ "name": <string>, "email": <string>, "password": <string> }`
    - Response: 200 OK on success, 400 Bad Request if the request format is invalid.
    - Curl: `curl -X POST -H "Content-Type: application/json" -d '{"name": "<name>", "email": "<email>", "password": "<password>"}' http://localhost:<port>/signup`

- **Login** - `POST /login`
    - Request Body: `{ "email": <string>, "password": <string> }`
    - Response: 200 OK on success (Returns JWT token), 400 Bad Request if the request format is invalid, 401 Unauthorized if username and password doesn't match.
    - Curl: `curl -X POST -H "Content-Type: application/json" -d '{"email": "<email>", "password": "<password>"}' http://localhost:<port>/login`

- **List Notes** - `GET /notes/list`
    - Headers: Authorization: Bearer <JWT token>
    - Response: 200 OK on success, Returns list of notes, 401 Unauthorized if JWT token is invalid.
    - Curl: `curl -X GET -H "Authorization: Bearer <JWT token>" http://localhost:<port>/notes/list`

- **Create Note** - `POST /notes/create`
    - Headers: Authorization: Bearer <JWT token>
    - Request Body: `{ "content": <string> }`
    - Response: 200 OK on success, 400 Bad Request if the request format is invalid, 401 Unauthorized if JWT token is invalid.
    - Curl: `curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer <JWT token>" -d '{"content": "<note content>"}' http://localhost:<port>/notes/create`

- **Delete Note** - `GET /notes/delete/:id`
    - Headers: Authorization: Bearer <JWT token>
    - Response: 200 OK on success, 400 Bad Request if the note ID is invalid, 401 Unauthorized if JWT token is invalid.
    - Curl: `curl -X GET -H "Authorization: Bearer <JWT token>" http://localhost:<port>/notes/delete/<note_id>`

Please replace `<string>`, `<JWT token>`, `<port>`, `<name>`, `<email>`, `<password>`, `<note content>`, and `<note_id>` with actual values. By default the application runs on 8080. 

## Author

sprectza
