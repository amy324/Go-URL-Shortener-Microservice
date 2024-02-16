

---

# Go URL Shortener Microservice

This is a URL shortener microservice implemented in Golang. It provides a RESTful API for shortening long URLs and redirecting to the original long URLs using short codes.

## Installation

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/your-username/Go-URL-Shortener-Microservice.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Go-URL-Shortener-Microservice
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Create a `.env` file in the root directory of the project. The file should contain the following environment variables:

    ```plaintext
    REDIS_STRING=your_redis_address:port
    REDIS_PASSWORD=your_redis_password
    ```

    Replace `your_redis_address:port` and `your_redis_password` with your actual Redis server address, port, and password.

## Usage

1. Run the application:

    ```bash
    go run .
    ```

2. The server should now be running on port `8080`.

3. Use the following endpoints to interact with the API:

    - `POST /shorten`: Shortens a long URL. Send a JSON object with the long URL in the request body.
    - `GET /{shortURL}`: Redirects to the original long URL associated with the provided short code.

## Testing

The project includes a test file named `main_test.go` in the root directory. This file contains unit tests for the main functionalities of the URL shortener microservice. To run the tests, use the following command:

```bash
go test
```

The tests ensure that the URL shortening and redirection functionalities work as expected and handle different scenarios correctly.

## Dependencies

This project relies on the following external dependencies:

- [gorilla/mux](https://github.com/gorilla/mux): A powerful HTTP router and URL matcher for building Go web servers.
- [joho/godotenv](https://github.com/joho/godotenv): A Go package for loading environment variables from a `.env` file.
- [go-redis/redis](https://github.com/go-redis/redis/v9): A Go client library for the Redis key-value store.

---
