# Drugstore Backend Application in Go

This project is a backend application for a drugstore implemented in Go. It provides core functionalities required to manage drugstore operations, including inventory management, order processing, and user authentication. The application is designed to be scalable and efficient, leveraging Go's concurrency model to handle multiple requests simultaneously.

## Features

- **`Inventory Management`**: Manage drug inventory, including adding, updating, and deleting drug records.
- **`Order Processing`**: Handle customer orders, including order creation, updating, and tracking.
- **`User Authentication`**: Secure user login and registration with encrypted passwords.
- **`RESTful API`**: Expose core functionalities through a RESTful API, enabling easy integration with front-end applications or other services.
- **`Concurrency Handling`**: Utilize Go's goroutines and channels for efficient request handling and background processing.

## Installation

1. **Clone the Repository**:

    ```bash
    git clone <URL-of-the-repository>
    ```

2. **Navigate to the Project Directory**:

    ```bash
    cd drugstore-backend
    ```

3. **Install Dependencies**:

    Ensure you have Go installed. Then run:

    ```bash
    go mod tidy
    ```

4. **Run the Application**:

    ```bash
    go run main.go
    ```

    The application will start on port `8080` by default.

## Configuration

- **Environment Variables**: Configure environment variables for database connections, API keys, and other settings. Create a `.env` file in the root directory and set the necessary variables.

    Example:

    ```env
    DATABASE_URL=your_database_url
    PORT=8080
    ```

- **Configuration File**: Modify `config.json` to adjust application settings such as logging level and API endpoints.

## Usage

- **API Endpoints**:

    - **GET /drugs**: Retrieve the list of drugs in inventory.
    - **POST /drugs**: Add a new drug to inventory.
    - **PUT /drugs/{id}**: Update drug details.
    - **DELETE /drugs/{id}**: Remove a drug from inventory.
    - **POST /orders**: Create a new order.
    - **GET /orders/{id}**: Get details of an order.

- **Authentication**: Use JWT tokens for secure access. Include the token in the `Authorization` header of your requests.

## Testing

- **Unit Tests**: Run unit tests using Go's testing framework:

    ```bash
    go test ./...
    ```

- **Integration Tests**: Ensure all components work together as expected. Refer to the `test` directory for integration test scripts.

## Contributing

1. **Fork the Repository**: Create a personal copy of the repository by forking it on GitHub.
2. **Create a Branch**: Make a new branch for your changes:

    ```bash
    git checkout -b feature/your-feature
    ```

3. **Make Changes**: Implement your changes and test them.
4. **Submit a Pull Request**: Push your branch and create a pull request on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
