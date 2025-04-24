# Go Web Service with Gin

This project demonstrates a simple RESTful web service built with Go and the Gin web framework. It provides a basic API for managing a collection of music albums.

## Project Overview

The application implements a REST API with the following endpoints:

- `GET /albums` - Retrieves all albums
- `GET /albums/:id` - Retrieves a specific album by ID
- `POST /albums` - Adds a new album

## Getting Started

### Prerequisites

- Go 1.24 or newer
- Git

### Installation

1. Clone the repository
2. Install dependencies:

```sh
go mod tidy
```

## Usage

### Running the server

Start the web server with:

```sh
go run main.go
```

The server will start on `localhost:8080`.

### API Endpoints

#### Get all albums

```
GET /albums
```

Response: Array of album objects

#### Get a specific album

```
GET /albums/:id
```

Response: Album object or 404 if not found

#### Add a new album

```
POST /albums
```

Request body:
```json
{
  "id": "4",
  "title": "Giant Steps",
  "artist": "John Coltrane", 
  "price": 63.99
}
```

Response: The created album with status 201

## Testing

The project includes comprehensive tests for all API endpoints. Run the tests with:

```sh
go test
```

The test suite covers:
- Retrieving all albums
- Retrieving a specific album by ID (both existing and non-existing IDs)
- Adding a new album with valid JSON data
- Error handling for invalid JSON requests

## Project Structure

- `main.go` - Core application with route definitions and handlers
- `main_test.go` - Test suite for API endpoints

## Dependencies

- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) - Web framework
- [github.com/stretchr/testify](https://github.com/stretchr/testify) - Testing toolkit