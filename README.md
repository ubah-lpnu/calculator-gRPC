# Caluculator-Protobuf

## Overview
This project implements a gRPC service for performing calculations such as multiplication and division of numbers. The service provides an API endpoint that accepts two numbers and returns the result of multiplying and dividing them.

## Features
- gRPC Service: Implements a gRPC service with a single API endpoint for performing calculations.
- Concurrent Operations: Utilizes goroutines for concurrent multiplication and division operations.
- Error Handling: Properly handles errors such as division by zero.
- Logging: Integrates logging using the zap logger for tracking application events.
- Unit Tests: Includes unit tests to ensure the correctness of the calculation logic.
- Docker: Provides a Dockerfile for containerizing the server application.

## Installation
To install the project, make sure you have Go installed on your machine. Then, run the following command:
```
git clone https://github.com/your_username/calculator-gRPC.git
cd calculator-gRPC
```

Install dependencies:
```
go mod download
```

## Usage
Start the server:
```
go run cmd/server/main.go
```
Run the client to perform calculations:
```
go run cmd/calculate_client/main.go -a=6.0 -b=3.0
```
## Tests
```
go test ./...
```
## Docker
Create docker image:
```
docker build -t calculator-server .
```
Run the server in docker container:
```
docker run -d -p 50051:50051 calculator-server
```

