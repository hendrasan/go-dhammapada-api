# Go Dhammapada API

## Description
The Go Dhammapada API is a RESTful service that provides access to the verses of the Dhammapada, a collection of sayings of the Buddha in verse form. This API allows users to retrieve verses, search for specific terms, and get information about the chapters.

## Installation
To install and run the project locally, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/hendrasan/go-dhammapada-api.git
    ```
2. Navigate to the project directory:
    ```sh
    cd go-dhammapada-api
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```
4. Copy .env.example to .env then update .env with your db (postgresql) credentials
    ```sh
    cp .env.example .env
    ```
5. Run the application:
    ```sh
    go run cmd/api/main.go
    ```

## Usage
Once the application is running, you can access the API at `http://localhost:8080`. Use a tool like `curl` or Postman to interact with the endpoints.

## Endpoints
Here are some of the available endpoints:

- `GET /health`: Health check

- `GET /api/v1/chapters`: Retrieve all chapters.
- `GET /api/v1/chapters/{id}`: Retrieve a specific chapter by ID.

- `GET /api/v1/verses`: Retrieve all verses.
- `GET /api/v1/verses/{id}`: Retrieve a specific verse by ID.
- `GET /api/v1/verses/random`: Retrieve a random verse.

- `GET /api/v1/search?q={term}`: Search for verses containing a specific term.