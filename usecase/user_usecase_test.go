package usecase_test

import (
	"context"
	"loan_tracker_api/domain"
	"loan_tracker_api/mocks"
	"loan_tracker_api/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// UserUseCasetestSuite struct to hold any shared resources or setup for the tests
type UserUseCasetestSuite struct {
	suite.Suite
	mockUserRepository *mocks.UserRepository
	UserUsecase        domain.UserUsecase
}

// SetupTest runs before each test case
func (s *UserUseCasetestSuite) SetupTest() {
	s.mockUserRepository = new(mocks.UserRepository)
	s.UserUsecase = usecase.NewUserUsecase(s.mockUserRepository, time.Second*2)
}

// TearDownTest runs after each test case
func (s *UserUseCasetestSuite) TearDownTest() {
	// Clean up resources if needed
}

// TestRegisterUser test the RegisterUser method
func (s *UserUseCasetestSuite) TestRegisterUser() {
	// Define the expected user data
	expectedUser := domain.User{
		UserName: "testuser",
		Email:    "test@gmail.com",
		Password: "passwoRd123!",
	}

	// Set up the mock expectation
	s.mockUserRepository.On("RegisterUser", &expectedUser).Return(nil).Once()

	// Call the method
	err := s.UserUsecase.RegisterUser(context.Background(), &expectedUser)

	// Check if the method returned an error
	s.NoError(err)
}

// TestVerifyUserEmail test the VerifyUserEmail method
func (s *UserUseCasetestSuite) TestVerifyUserEmail() {
	// Define the expected token
	token := "testtoken"

	// Set up the mock expectation
	s.mockUserRepository.On("VerifyUserEmail", token).Return(nil).Once()

	// Call the method
	err := s.UserUsecase.VerifyUserEmail(context.Background(), token)

	// Check if the method returned an error
	s.NoError(err)
}

// TestLoginUser test the LoginUser method
func (s *UserUseCasetestSuite) TestLoginUser() {
	// Define the expected user data
	expectedUser := domain.User{
		Email:    "test@gmail.com",
		Password: "passwoRd123!",
	}

	// Set up the mock expectation
	s.mockUserRepository.On("LoginUser", expectedUser).Return("token", "anothertoken", nil).Once()

	// Call the method
	_, _, err := s.UserUsecase.LoginUser(context.Background(), expectedUser)

	// Check if the method returned an error
	s.NoError(err)
}

// TestTokenRefresh test the TokenRefresh method
func (s *UserUseCasetestSuite) TestTokenRefresh() {
	// Define the expected token
	refreshToken := "testtoken"

	// Set up the mock expectation
	s.mockUserRepository.On("TokenRefresh", refreshToken).Return("token", nil).Once()

	// Call the method
	_, err := s.UserUsecase.TokenRefresh(context.Background(), refreshToken)

	// Check if the method returned an error
	s.NoError(err)
}

// TestUserProfile test the UserProfile method
func (s *UserUseCasetestSuite) TestUserProfile() {
	// Define the expected user ID
	uid := "testuid"

	// Set up the mock expectation
	s.mockUserRepository.On("UserProfile", uid).Return(domain.User{}, nil).Once()

	// Call the method
	_, err := s.UserUsecase.UserProfile(context.Background(), uid)

	// Check if the method returned an error
	s.NoError(err)
}

// TestForgotPassword test the ForgotPassword method
func (s *UserUseCasetestSuite) TestForgotPassword() {
	// Define the expected email
	email := "test@gmail.com"

	// Set up the mock expectation
	s.mockUserRepository.On("ForgotPassword", email).Return(nil).Once()

	// Call the method
	err := s.UserUsecase.ForgotPassword(context.Background(), email)

	// Check if the method returned an error
	s.NoError(err)
}

// TestResetPassword test the ResetPassword method
func (s *UserUseCasetestSuite) TestResetPassword() {
	token := "test-token"
	newpass := "password123"

	//setup mock expectations
	s.mockUserRepository.On("ResetPassword", token, newpass).Return(nil).Once()

	//call the method
	err := s.UserUsecase.ResetPassword(context.Background(), token, newpass)

	//check if the method returned an error
	s.NoError(err)
}

// TestUpdateUserDetails test the UpdateUserDetails method
func (s *UserUseCasetestSuite) TestUpdateUserDetails() {
	// Define the expected user data
	expectedUser := domain.User{
		UserName: "testuser",
		Email:    "test@gmail.com",
	}

	// Set up the mock expectation
	s.mockUserRepository.On("UpdateUserDetails", &expectedUser).Return(nil).Once()

	// Call the method
	err := s.UserUsecase.UpdateUserDetails(context.Background(), &expectedUser)

	// Check if the method returned an error
	s.NoError(err)

}

// TestLogoutUser test the LogoutUser method
func (s *UserUseCasetestSuite) TestLogoutUser() {
	// Define the expected user ID
	uid := "testuid"

	// Set up the mock expectation
	s.mockUserRepository.On("LogoutUser", uid).Return(nil).Once()

	// Call the method
	err := s.UserUsecase.LogoutUser(context.Background(), uid)

	// Check if the method returned an error
	s.NoError(err)
}

// TestViewAllUsers test the ViewAllUsers method
func (s *UserUseCasetestSuite) TestViewAllUsers() {
	// Set up the mock expectation
	s.mockUserRepository.On("ViewAllUsers").Return([]domain.User{}, nil).Once()

	// Call the method
	_, err := s.UserUsecase.ViewAllUsers(context.Background())

	// Check if the method returned an error
	s.NoError(err)
}

// Run the test suite
func TestUserUsecaseRunSuite(t *testing.T) {
	suite.Run(t, new(UserUseCasetestSuite))
}
