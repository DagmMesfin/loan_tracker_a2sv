# Loan Tracker API - Project Documentation

## Introduction

The Loan Tracker API is a backend system designed to manage loan applications and user accounts within a financial service context. Built using Golang with the Gin framework and adhering to clean architecture principles, this API offers a robust and scalable solution for handling various functionalities, including user management, loan processing, and administrative tasks. The API ensures secure and efficient data handling, making it a reliable foundation for financial services.

This documentation provides an overview of the API's functionalities, including endpoints for user and admin operations, authentication mechanisms, and system settings. The Loan Tracker API is designed to be flexible and extendable, supporting the evolving needs of a financial service provider.

## Technologies Used

- **Go**: A statically typed, compiled programming language known for its simplicity and performance.
- **Gin Framework**: A lightweight and fast HTTP web framework for Go, ideal for building RESTful APIs.
- **MongoDB**: A NoSQL database known for its flexibility and scalability, used for storing task and user data.
- **JWT (JSON Web Tokens)**: Used for securing API endpoints by providing token-based authentication.

## The Endpoints

### 1. User Management

#### 1.1 User Registration
- **Endpoint:** `POST /users/register`
- **Description:** This endpoint allows new users to register by providing their email, password, and profile details. Upon successful registration, the user receives a confirmation message.
- **Response:** Success or error message.

#### 1.2 Email Verification
- **Endpoint:** `GET /users/verify-email`
- **Description:** After registration, users must verify their email addresses. The system sends a unique verification token to the user's email, which must be submitted to this endpoint to activate the account.
- **Parameters:**
  - `token`: Verification token sent via email.
  - `email`: User's email address.
- **Response:** Success or error message.

#### 1.3 User Login
- **Endpoint:** `POST /users/login`
- **Description:** This endpoint authenticates users by validating their credentials and, if successful, provides access and refresh tokens for future requests.
- **Response:** Success message with access and refresh tokens, or error message.

#### 1.4 Token Refresh
- **Endpoint:** `POST /users/token/refresh`
- **Description:** This endpoint allows users to refresh their access token using a valid refresh token, ensuring continued access without needing to reauthenticate.
- **Response:** New access token or error message.

#### 1.5 User Profile
- **Endpoint:** `GET /users/profile`
- **Description:** This endpoint retrieves the profile information of the authenticated user.
- **Response:** User profile data.

#### 1.6 Password Reset Request
- **Endpoint:** `POST /users/password-reset`
- **Description:** Users can request a password reset by providing their email. The system sends a password reset link to the provided email address.
- **Response:** Success or error message.

#### 1.7 Password Update After Reset
- **Endpoint:** `POST /users/password-update`
- **Description:** After receiving the password reset link, users can update their password by submitting the new password and the token received via email.
- **Response:** Success or error message.

### 2. Admin Functionalities

#### 2.1 View All Users
- **Endpoint:** `GET /admin/users`
- **Description:** This endpoint allows admins to retrieve a list of all registered users.
- **Response:** A list of users or an error message.

#### 2.2 Delete User Account
- **Endpoint:** `DELETE /admin/users/{id}`
- **Description:** Admins can delete a specific user account by providing the user's ID.
- **Response:** Success or error message.

## Running the API

1. **Install Dependencies:** Ensure you have Go and MongoDB installed.
2. **Clone the Repository:**
   ```bash
   git clone https://github.com/DagmMesfin/loan_tracker_a2sv.git
   ```
3. **Navigate to the Project Directory:**
   ```bash
   cd loan_tracker_a2sv
   ```
4. **Set Environment Variables:**
   ```bash
   export JWT_SECRET=<your_jwt_secret>
   export MONGO_URI=<your_mongo_uri>
   ```
5. **Run the Application:**
   ```bash
   go run main.go
   ```

