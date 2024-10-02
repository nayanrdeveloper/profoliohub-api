# ProFolioHub API

This project is a Dynamic portfolio service built using **GoLang**, **Gin Web Framework**, and **MongoDB**.

## Project Structure

/go-user-auth │ ├── /config # Configuration settings like env variables ├── /controllers # Logic for API controllers ├── /models # Database models ├── /routes # API routes ├── /services # Services for handling database interactions ├── /middlewares # Middleware for handling authentication, logging, etc. ├── /helpers # Reusable helper functions (e.g., JWT, hashing, etc.) ├── /database # MongoDB connection setup ├── /docs # API documentation (e.g., Swagger) ├── .env # Environment variables (DB connection, secret keys, etc.) ├── main.go # Entry point for the application └── go.mod # Go module configuration

## Requirements

- Go 1.16+
- MongoDB 4.0+
- Gin Web Framework
- JWT for authentication
- bcrypt for password hashing
- Godotenv for environment variable management

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/go-user-auth.git
    cd go-user-auth
    ```

2. Install the required Go modules:

    ```bash
    go mod tidy
    ```

3. Create an `.env` file in the root directory with the following content:

    ```env
    PORT=8080
    MONGO_URI=mongodb://localhost:27017
    JWT_SECRET=your_jwt_secret
    ```

4. Start the application:

    ```bash
    go run main.go
    ```