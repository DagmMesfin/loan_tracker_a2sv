Certainly! Here's the Loan Tracker API documentation based on your provided template and router endpoints:

---

# Loan Tracker API Documentation

## Overview
The Loan Tracker API is a robust and scalable solution designed to manage loan tracking and related functionalities for both users and administrators. Built with Go, leveraging the Gin framework, this API adheres to Clean Architecture principles, ensuring maintainability, testability, and scalability.

## Core Features

### User Management
- **Registration and Login**: Secure authentication with password hashing and JWT token issuance.
- **Email Verification**: Users can verify their email addresses after registration.
- **Profile Management**: Users can view and update their profiles.
- **Password Reset**: Users can reset their password securely via email.
- **Loan Applications**: Users can apply for loans, view their loan details, and track repayments.
- **Token Management**: Support for token refresh and user logout.

### Admin Management
- **User Management**: Admins can manage user accounts, including viewing all users and deleting user accounts.
- **Loan Management**: Admins can review, approve, or reject loan applications, manage loan details, and delete loans.
- **System Logs**: Admins can view system logs for audit purposes.

## Clean Architecture Implementation
The API is structured according to the Clean Architecture paradigm, which separates the code into distinct layers, ensuring that business rules are isolated from implementation details:

- **Domain Layer**: Contains core business logic and entities. This is the heart of the application, encapsulating business rules that are independent of external systems.

- **Use Case Layer**: Implements application-specific business rules. It orchestrates the flow of data to and from the entities, ensuring that all the business rules are followed.

- **Interface Adapters**: Includes controllers, gateways, and presenters. This layer translates data between the use case layer and the external systems, like the web framework (Gin), databases, and third-party APIs.

- **Frameworks and Drivers**: Contains the details of frameworks, databases, and other implementation details. This layer is responsible for implementing the details specified by the use cases.

## Authentication and Security
The API uses JWT (JSON Web Tokens) for securing endpoints. Access tokens are issued upon successful authentication and must be provided in the Authorization header for protected routes. Refresh tokens are used to obtain new access tokens when the original tokens expire.

## Routes and Endpoints

### User Routes
- **POST /user/register**: Register a new user.
- **POST /user/verify-email**: Verify a user's email address.
- **POST /user/login**: Login and receive an access token.
- **GET /user/token-refresh**: Refresh the access token (requires authentication).
- **GET /user/profile**: Retrieve user profile information (requires authentication).
- **GET /user/logout**: Log out the user (requires authentication).
- **PUT /user/update**: Update user profile information (requires authentication).
- **POST /user/password-reset**: Initiate a password reset.
- **POST /user/password-update**: Update the password after a reset.

### Loan Routes
- **POST /loan/apply**: Submit a loan application (requires authentication).
- **GET /loan/:loan_id**: View loan details by ID (requires authentication).

### Admin Routes
- **GET /admin/users**: List all users (requires admin authentication).
- **DELETE /admin/user/:id**: Delete a user by ID (requires admin authentication).
- **GET /admin/loans**: List all loan applications (requires admin authentication).
- **PATCH /admin/loans/:loan_id/status**: Approve or reject a loan application by ID (requires admin authentication).
- **DELETE /admin/loans/:loan_id**: Delete a loan by ID (requires admin authentication).
- **GET /admin/logs**: View system logs (requires admin authentication).

## Testing and Validation
The API includes comprehensive unit tests to validate business logic at the domain and use case layers, ensuring that all critical functionalities work as expected. Integration tests are also implemented to validate the interaction between different layers of the application.

## Postman Collection
A Postman collection is available for easy testing of the API endpoints. The collection includes pre-configured requests for all major routes, ensuring you can test the API quickly and efficiently.

## Repository Structure
The repository follows a clear and organized structure to support Clean Architecture:

- **/cmd**: Entry point of the application.
- **/internal**:
  - **/domain**: Business entities and logic.
  - **/usecase**: Application business rules.
  - **/interfaces**: Adapters for the web framework, databases, etc.
  - **/frameworks**: Implementation of frameworks and drivers.
  - **/tests**: Unit and integration tests.

## Installation and Setup

1. **Clone the repository:**
    ```bash
    git clone <repository-url>
    cd loan-tracker-api
    ```

2. **Install dependencies:**
    ```bash
    go mod tidy
    ```

3. **Run the application:**
    ```bash
    go run cmd/main.go
    ```

4. **Run tests:**
    ```bash
    go test ./...
    ```

---

for more information, go here: https://documenter.getpostman.com/view/37336034/2sAXjGduYB

This documentation provides a comprehensive overview of the Loan Tracker API, covering all key aspects, from core features to implementation and testing.