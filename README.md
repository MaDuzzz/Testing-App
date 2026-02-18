# Testing App - Go Backend

A simple Go backend application using Gin framework with a single GET API endpoint.

## Prerequisites

- Go 1.21 or higher

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Build the application:
```bash
go build -o testing-app
```

3. Run the application:
```bash
./testing-app
```

The server will start on `http://localhost:8080`

## API Endpoint

### GET /api/hello

Returns a JSON response with a greeting message.

**Response:**
```json
{
  "message": "Hello from Go Backend with Gin!",
  "status": "success"
}
```

**Example:**
```bash
curl http://localhost:8080/api/hello
```