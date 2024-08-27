package controllers_test

import (
	"bytes"
	"loan_tracker_api/deliveries/controllers"
	"loan_tracker_api/domain"
	"loan_tracker_api/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
	controller  *controllers.UserController
	mockUsecase *mocks.UserUsecase
	Recorder    *httptest.ResponseRecorder
	mockContext *gin.Context
}

func (suite *UserControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockUsecase = new(mocks.UserUsecase)
	suite.controller = controllers.NewUserController(suite.mockUsecase)
	// Prepare the recorder and context
	suite.Recorder = httptest.NewRecorder()
	suite.mockContext, _ = gin.CreateTestContext(suite.Recorder)
}

func (suite *UserControllerTestSuite) TestRegisterUser() {
	// Define the expected user data
	expectedUser := domain.User{
		UserName: "testuser",
		Email:    "test@example.com",
		Password: "passwoRd123!",
	}

	// Set up the mock expectation
	suite.mockUsecase.On("RegisterUser", mock.Anything, mock.MatchedBy(func(user *domain.User) bool {
		return user.UserName == expectedUser.UserName &&
			user.Email == expectedUser.Email &&
			user.Password == expectedUser.Password
	})).Return(nil).Once()

	// Prepare the request body
	requestBody := `{
		"username": "testuser",
		"email": "test@example.com",
		"password": "passwoRd123!"
	}`

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("POST", "/user/register", strings.NewReader(requestBody))
	suite.mockContext.Request.Header.Set("Content-Type", "application/json")

	// Call the controller method
	suite.controller.RegisterUser(suite.mockContext)

	// Check the response
	suite.Equal(200, suite.Recorder.Code)
	suite.Contains(suite.Recorder.Body.String(), "User registered successfully")

}

func (suite *UserControllerTestSuite) TestVerifyEmail() {
	suite.mockUsecase.On("VerifyUserEmail", mock.Anything, mock.Anything).Return(nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("GET", "/user/verify-email?token=test-token", nil)

	// Call the controller method
	suite.controller.VerifyEmail(suite.mockContext)

	// Check the response
	suite.Equal(200, suite.Recorder.Code)
	suite.Contains(suite.Recorder.Body.String(), "Email verified successfully")
}

func (suite *UserControllerTestSuite) TestLoginUser() {
	user := domain.User{Email: "test@example.com", Password: "password123"}
	suite.mockUsecase.On("LoginUser", mock.Anything, user).Return("mocked-refresh-token", "mocked-access-token", nil).Once()

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = httptest.NewRequest(http.MethodPost, "/user/login", bytes.NewReader([]byte(`{"email":"test@example.com","password":"password123"}`)))

	suite.controller.LoginUser(context)

	suite.Equal(http.StatusOK, recorder.Code)
}

func (suite *UserControllerTestSuite) TestTokenRefresh() {
	suite.mockUsecase.On("TokenRefresh", mock.Anything, mock.Anything).Return("mocked-access-token", nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("GET", "/user/token-refresh?refresh-token=test-token", nil)

	// Call the controller method
	suite.controller.TokenRefresh(suite.mockContext)

	// Check the response
	suite.Equal(200, suite.Recorder.Code)
	suite.Contains(suite.Recorder.Body.String(), "{\"message\":\"token refreshed\",\"new-access-token\":\"mocked-access-token\"}")
}

func (suite *UserControllerTestSuite) TestUserProfile() {
	suite.mockUsecase.On("UserProfile", mock.Anything, mock.Anything).Return(domain.User{}, nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("GET", "/user/profile", nil)
	suite.mockContext.Set("userid", "test-user-id")

	// Call the controller method
	suite.controller.UserProfile(suite.mockContext)

	// Check the response
	suite.Equal(200, suite.Recorder.Code)

}

func (suite *UserControllerTestSuite) TestForgotPassword() {
	suite.mockUsecase.On("ForgotPassword", mock.Anything, mock.Anything).Return(nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("POST", "/user/password-reset", strings.NewReader(`{"email":"testmail@gmail.com"}`))
	suite.mockContext.Request.Header.Set("Content-Type", "application/json")

	// Call the controller method
	suite.controller.ForgotPassword(suite.mockContext)

	// Check the response
	suite.Equal(202, suite.Recorder.Code)
	suite.Contains(suite.Recorder.Body.String(), "{\"message\":\"email succefully sent to the email provided\"}")

}

func (suite *UserControllerTestSuite) TestResetPassword() {
	suite.mockUsecase.On("ResetPassword", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("POST", "/user/password-update?token=test-token", strings.NewReader(`{"password":"newPassword123!"}`))
	suite.mockContext.Request.Header.Set("Content-Type", "application/json")

	// Call the controller method
	suite.controller.ResetPassword(suite.mockContext)

	// Check the response
	suite.Equal(202, suite.Recorder.Code)
	suite.Contains(suite.Recorder.Body.String(), "{\"message\":\"Password has been reset successfully\"}")
}

func (suite *UserControllerTestSuite) TestUpdateUserDetails() {
	suite.mockUsecase.On("UpdateUserDetails", mock.Anything, mock.Anything).Return(nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("PUT", "/user/update", strings.NewReader(`{"username":"newusername"}`))
	suite.mockContext.Request.Header.Set("Content-Type", "application/json")
	suite.mockContext.Set("userid", "test-user-id")

	// Call the controller method
	suite.controller.UpdateUserDetails(suite.mockContext)

	// Check the response
	suite.Equal(200, suite.Recorder.Code)
	suite.Contains(suite.Recorder.Body.String(), "{\"message\":\"User details updated successfully\"}")

}

func (suite *UserControllerTestSuite) TestLogoutUser() {
	suite.mockUsecase.On("LogoutUser", mock.Anything, mock.Anything).Return(nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("GET", "/user/logout", nil)
	suite.mockContext.Set("userid", "test-user-id")

	// Call the controller method
	suite.controller.LogoutUser(suite.mockContext)

	// Check the response
	suite.Equal(200, suite.Recorder.Code)
	suite.Contains(suite.Recorder.Body.String(), "{\"message\":\"User logged out successfully\"}")
}

func (suite *UserControllerTestSuite) TestViewAllUsers() {
	suite.mockUsecase.On("ViewAllUsers", mock.Anything).Return([]domain.User{}, nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("GET", "/admin/users", nil)

	// Call the controller method
	suite.controller.ViewAllUsers(suite.mockContext)

	// Check the response
	suite.Equal(200, suite.Recorder.Code)

}

func (suite *UserControllerTestSuite) TestDeleteUser() {
	suite.mockUsecase.On("DeleteUser", mock.Anything, mock.Anything).Return(nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("DELETE", "/admin/user/test-user-id", nil)

	// Call the controller method
	suite.controller.DeleteUser(suite.mockContext)

	// Check the response
	suite.Equal(200, suite.Recorder.Code)
	suite.Contains(suite.Recorder.Body.String(), "{\"message\":\"User deleted successfully\"}")
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
