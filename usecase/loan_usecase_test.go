package usecase_test

import (
	"context"
	"loan_tracker_api/domain"
	"loan_tracker_api/mocks"
	"loan_tracker_api/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanUsecaseTestSuite struct {
	suite.Suite
	mockLoanRepository *mocks.LoanRepository
	LoanUsecase        domain.LoanUsecase
}

func (s *LoanUsecaseTestSuite) SetupTest() {
	s.mockLoanRepository = new(mocks.LoanRepository)
	s.LoanUsecase = usecase.NewLoanUsecase(s.mockLoanRepository, time.Second*2)
}

func (s *LoanUsecaseTestSuite) TearDownTest() {
	// Clean up resources if needed
}

func (s *LoanUsecaseTestSuite) TestApplyForLoan() {
	expectedLoan := domain.Loan{
		Amount:   100000,
		Duration: 12,
	}

	s.mockLoanRepository.On("ApplyForLoan", &expectedLoan, "testuserid").Return(nil).Once()

	err := s.LoanUsecase.ApplyForLoan(context.Background(), &expectedLoan, "testuserid")

	s.NoError(err)
}

func (s *LoanUsecaseTestSuite) TestLoanDetails() {
	expectedLoan := domain.Loan{
		ID: primitive.NewObjectID(),
	}

	s.mockLoanRepository.On("LoanDetails", "testloanid", "testuserid").Return(expectedLoan, nil).Once()

	loan, err := s.LoanUsecase.LoanDetails(context.Background(), "testloanid", "testuserid")

	s.NoError(err)
	s.Equal(expectedLoan, loan)
}

func (s *LoanUsecaseTestSuite) TestViewAllLoans() {
	expectedLoans := []domain.Loan{
		{
			ID: primitive.NewObjectID(),
		},
	}

	s.mockLoanRepository.On("ViewAllLoans", 1, "pending", "asc").Return(expectedLoans, 1, nil).Once()

	loans, _, err := s.LoanUsecase.ViewAllLoans(context.Background(), 1, "pending", "asc")

	s.NoError(err)
	s.Equal(expectedLoans, loans)
}

func (s *LoanUsecaseTestSuite) TestApproveRejectLoan() {
	s.mockLoanRepository.On("ApproveRejectLoan", "testloanid", "approved", "testuserid").Return(nil).Once()

	err := s.LoanUsecase.ApproveRejectLoan(context.Background(), "testloanid", "approved", "testuserid")

	s.NoError(err)
}

func (s *LoanUsecaseTestSuite) TestDeleteLoan() {
	s.mockLoanRepository.On("DeleteLoan", "testloanid", "testuserid").Return(nil).Once()

	err := s.LoanUsecase.DeleteLoan(context.Background(), "testloanid", "testuserid")

	s.NoError(err)
}

func (s *LoanUsecaseTestSuite) TestViewLogs() {
	expectedLogs := []domain.Log{
		{
			ID: primitive.NewObjectID(),
		},
	}

	s.mockLoanRepository.On("ViewLogs").Return(expectedLogs, nil).Once()

	logs, err := s.LoanUsecase.ViewLogs(context.Background())

	s.NoError(err)
	s.Equal(expectedLogs, logs)
}

func TestLoanUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LoanUsecaseTestSuite))
}
