package usecase

import (
	"context"
	"loan_tracker_api/domain"
	"time"
)

type LoanUsecase struct {
	UserRepo       domain.LoanRepository
	contextTimeout time.Duration
}

func NewLoanUsecase(Userrepo domain.LoanRepository, timeout time.Duration) domain.LoanUsecase {
	return &LoanUsecase{
		UserRepo:       Userrepo,
		contextTimeout: timeout,
	}

}

func (luse *LoanUsecase) ApplyForLoan(c context.Context, loan *domain.Loan, userid string) error {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.UserRepo.ApplyForLoan(loan, userid)
}

func (luse *LoanUsecase) LoanDetails(c context.Context, loanID string, userid string) (domain.Loan, error) {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.UserRepo.LoanDetails(loanID, userid)
}

func (luse *LoanUsecase) ViewAllLoans(c context.Context, pgnum int, status, order string) ([]domain.Loan, int, error) {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.UserRepo.ViewAllLoans(pgnum, status, order)
}

func (luse *LoanUsecase) ApproveRejectLoan(c context.Context, loanID string, status string, userid string) error {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.UserRepo.ApproveRejectLoan(loanID, status, userid)
}

func (luse *LoanUsecase) DeleteLoan(c context.Context, loanID string, userid string) error {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.UserRepo.DeleteLoan(loanID, userid)
}

func (luse *LoanUsecase) ViewLogs(c context.Context) ([]domain.Log, error) {
	_, cancel := context.WithTimeout(c, luse.contextTimeout)
	defer cancel()
	return luse.UserRepo.ViewLogs()
}
