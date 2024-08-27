package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Loan struct represents the loan model
type Loan struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Amount    float64            `json:"amount" bson:"amount"`
	Interest  float64            `json:"interest" bson:"interest"`
	Duration  int                `json:"duration" bson:"duration"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// LoanRepository represents the loan repository contract
type LoanRepository interface {
	ApplyForLoan(loan *Loan, userid string) error
	LoanDetails(loanID string, userid string) (Loan, error)
	ViewAllLoans(pgnum int, status, order string) ([]Loan, int, error)
	ApproveRejectLoan(loanID string, status, userid string) error
	DeleteLoan(loanID string, userid string) error
	ViewLogs() ([]Log, error)
}

// LoanUsecase represents the loan usecase contract
type LoanUsecase interface {
	ApplyForLoan(c context.Context, loan *Loan, userid string) error
	LoanDetails(c context.Context, loanID string, userid string) (Loan, error)
	ViewAllLoans(c context.Context, pgnum int, status, order string) ([]Loan, int, error)
	ApproveRejectLoan(c context.Context, loanID string, status, userid string) error
	DeleteLoan(c context.Context, loanID string, userid string) error
	ViewLogs(c context.Context) ([]Log, error)
}
