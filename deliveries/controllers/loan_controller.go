package controllers

import (
	"context"
	"loan_tracker_api/domain"
	"net/http"
	"strconv"

	gin "github.com/gin-gonic/gin"
)

// LoanController struct to hold the usecase
type LoanController struct {
	LoanUsecase domain.LoanUsecase
}

// NewLoanController function to create a new LoanController
func NewLoanController(luse domain.LoanUsecase) *LoanController {
	return &LoanController{
		LoanUsecase: luse,
	}
}

// ApplyForLoan function to handle the ApplyForLoan endpoint
func (lc *LoanController) ApplyForLoan(c *gin.Context) {
	userid := c.GetString("userid")
	var loan domain.Loan

	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	err := lc.LoanUsecase.ApplyForLoan(context.Background(), &loan, userid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Loan application successful", "loan": loan})
}

// LoanDetails function to handle the LoanDetails endpoint
func (lc *LoanController) LoanDetails(c *gin.Context) {
	userid := c.GetString("userid")
	loanID := c.Param("loan_id")

	loan, err := lc.LoanUsecase.LoanDetails(context.Background(), loanID, userid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"loan": loan})
}

// ViewAllLoans function to handle the ViewAllLoans endpoint
func (lc *LoanController) ViewAllLoans(c *gin.Context) {
	pgnum, err := strconv.Atoi(c.Query("pgnum"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	status := c.Query("status")
	order := c.Query("order")

	loans, _, err := lc.LoanUsecase.ViewAllLoans(context.Background(), pgnum, status, order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"loans": loans})
}

// ApproveRejectLoan function to handle the ApproveRejectLoan endpoint
func (lc *LoanController) ApproveRejectLoan(c *gin.Context) {
	userid := c.GetString("userid")
	loanID := c.Param("loan_id")

	var status struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if status.Status != "approved" && status.Status != "rejected" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	err := lc.LoanUsecase.ApproveRejectLoan(context.Background(), loanID, status.Status, userid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan status updated"})
}

// DeleteLoan function to handle the DeleteLoan endpoint
func (lc *LoanController) DeleteLoan(c *gin.Context) {
	userid := c.GetString("userid")
	loanID := c.Param("loan_id")

	err := lc.LoanUsecase.DeleteLoan(context.Background(), loanID, userid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted"})
}

// ViewLogs function to handle the ViewLogs endpoint
func (lc *LoanController) ViewLogs(c *gin.Context) {
	logs, err := lc.LoanUsecase.ViewLogs(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"logs": logs})
}
