# swift-api
This is the repository for the swift-api home exercise.

## Swift API

### How to run the application:

1. **Locally**:
   1. Install Go (version 1.21 or newer).
   2. Clone the repository:
      ```bash
      git clone https://github.com/yourusername/swift-api.git
      cd swift-api
      ```
   3. Install dependencies:
      ```bash
      go mod tidy
      ```
   4. Run the command:
      ```bash
      go run main.go
      ```
   5. The application will be available at `http://localhost:8080`.

2. **With Docker**:
   1. Install Docker and Docker Compose.
   2. Clone the repository:
      ```bash
      git clone https://github.com/yourusername/swift-api.git
      cd swift-api
      ```
   3. Run the application and database in containers:
      ```bash
      docker-compose up --build
      ```
   4. The application will be available at `http://localhost:8080`.

### How to test the API:

Example cURL requests:

- **GET**: Get SWIFT code data:
  ```bash
  curl http://localhost:8080/v1/swift-codes/XXXXX

POST: Add a new SWIFT code:

curl -X POST -d '{"swiftCode": "XXXXX", "bankName": "Bank", "address": "Address", "countryISO2": "US", "countryName": "USA", "isHeadquarter": true}' http://localhost:8080/v1/swift-codes

DELETE: Delete a SWIFT code:

curl -X DELETE http://localhost:8080/v1/swift-codes/XXXXX

###Example responses:

#GET:

{
  "swiftCode": "XXXXX",
  "bankName": "Bank",
  "address": "Address",
  "countryISO2": "US",
  "countryName": "USA",
  "isHeadquarter": true
}

#POST: Response on success:

{
  "message": "Swift code successfully added"
}

#DELETE: Response on success:

{
  "message": "Swift code successfully deleted"
}

###Unit testing

Unit tests for business logic and endpoints are included in the repository. You can run them using the Go testing tool, for example, go test:

go test ./...

###Dependencies

Go 1.21 or newer.

PostgreSQL (or Docker with a PostgreSQL container).

###Summary

The application allows managing SWIFT code data via an API. You can run it locally or in a Docker container and test various endpoints using cURL or Postman.

###What's included:

Docker installation and configuration – details on running the application and database in containers.

More details on unit testing – how to run tests in Go.

API response examples – the user knows what to expect after making a request.

Instructions on dependencies – Go 1.21, PostgreSQL, and Docker.