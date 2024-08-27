package controllers_test

import (
	"loan_tracker_api/deliveries/controllers"
	"loan_tracker_api/domain"
	"loan_tracker_api/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gin "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanControllerTestSuite struct {
	suite.Suite
	controller  *controllers.LoanController
	mockUsecase *mocks.LoanUsecase
	Recorder    *httptest.ResponseRecorder
	mockContext *gin.Context
}

func (suite *LoanControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockUsecase = new(mocks.LoanUsecase)
	suite.controller = controllers.NewLoanController(suite.mockUsecase)
	// Prepare the recorder and context
	suite.Recorder = httptest.NewRecorder()
	suite.mockContext, _ = gin.CreateTestContext(suite.Recorder)
}

func (suite *LoanControllerTestSuite) TestApplyForLoan() {

	// Set up the mock expectation
	suite.mockUsecase.On("ApplyForLoan", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

	// Prepare the request body
	requestBody := `{
		"amount": 100000,
		"duration": 12
	}`

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("POST", "/loan/apply", strings.NewReader(requestBody))
	suite.mockContext.Request.Header.Set("Content-Type", "application/json")

	// Call the controller function
	suite.controller.ApplyForLoan(suite.mockContext)

	// Check the response
	suite.Equal(http.StatusCreated, suite.Recorder.Code)
}

func (suite *LoanControllerTestSuite) TestLoanDetails() {
	// Define the expected loan data
	expectedLoan := domain.Loan{
		ID: primitive.NewObjectID(),
	}

	// Set up the mock expectation
	suite.mockUsecase.On("LoanDetails", mock.Anything, mock.Anything, mock.Anything).Return(expectedLoan, nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("GET", "/loan/testloanid", nil)

	// Call the controller function
	suite.controller.LoanDetails(suite.mockContext)

	// Check the response
	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func (suite *LoanControllerTestSuite) TestViewAllLoans() {
	// Define the expected loans data
	expectedLoans := []domain.Loan{
		{
			ID: primitive.NewObjectID(),
		},
	}

	// Set up the mock expectation
	suite.mockUsecase.On("ViewAllLoans", mock.Anything, 1, "pending", "asc").Return(expectedLoans, 1, nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("GET", "/loans", nil)
	suite.mockContext.Request.URL.RawQuery = "pgnum=1&status=pending&order=asc"

	// Call the controller function
	suite.controller.ViewAllLoans(suite.mockContext)

	// Check the response
	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func (suite *LoanControllerTestSuite) TestApproveRejectLoan() {
	// Set up the mock expectation
	suite.mockUsecase.On("ApproveRejectLoan", mock.Anything, "testloanid", "approved", mock.Anything).Return(nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("PUT", "/loan/testloanid/approve", strings.NewReader(`{"status": "approved"}`))
	suite.mockContext.Params = append(suite.mockContext.Params, gin.Param{Key: "loan_id", Value: "testloanid"})
	suite.mockContext.Set("userid", "testuserid")
	suite.mockContext.Request.Header.Set("Content-Type", "application/json")

	// Call the controller function
	suite.controller.ApproveRejectLoan(suite.mockContext)

	// Check the response
	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func (suite *LoanControllerTestSuite) TestDeleteLoan() {
	// Set up the mock expectation
	suite.mockUsecase.On("DeleteLoan", mock.Anything, "testloanid", mock.Anything).Return(nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("DELETE", "/loan/testloanid", nil)
	suite.mockContext.Params = append(suite.mockContext.Params, gin.Param{Key: "loan_id", Value: "testloanid"})
	suite.mockContext.Set("userid", "testuserid")

	// Call the controller function
	suite.controller.DeleteLoan(suite.mockContext)

	// Check the response
	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func (suite *LoanControllerTestSuite) TestViewLogs() {
	// Define the expected logs data
	expectedLogs := []domain.Log{
		{
			ID: primitive.NewObjectID(),
		},
	}

	// Set up the mock expectation
	suite.mockUsecase.On("ViewLogs", mock.Anything).Return(expectedLogs, nil).Once()

	// Prepare the request
	suite.mockContext.Request = httptest.NewRequest("GET", "/logs", nil)

	// Call the controller function
	suite.controller.ViewLogs(suite.mockContext)

	// Check the response
	suite.Equal(http.StatusOK, suite.Recorder.Code)
}

func TestLoanControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LoanControllerTestSuite))
}
