package controllers

import (
	"context"
	"loan_tracker_api/domain"
	"net/http"
	"strconv"

	gin "github.com/gin-gonic/gin"
)

type LoanController struct {
	LoanUsecase domain.LoanUsecase
}

func NewLoanController(luse domain.LoanUsecase) *LoanController {
	return &LoanController{
		LoanUsecase: luse,
	}
}

// ApplyForLoan godoc
// @Summary Apply for a loan
// @Description Apply for a loan
// @Tags Loan
// @Accept json
// @Produce json
// @Param loan body domain.Loan true "Loan details"
// @Success 201 {object} domain.Loan
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /loan/apply [post]
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

// LoanDetails godocs
// @Summary Get loan details
// @Description Get loan details
// @Tags Loan
// @Accept json
// @Produce json
// @Param loan_id path string true "Loan ID"
// @Success 200 {object} domain.Loan
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /loan/{loan_id} [get]
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

// ViewAllLoans godoc
// @Summary View all loans
// @Description View all loans
// @Tags Admin
// @Accept json
// @Produce json
// @Param pgnum query int false "Page number"
// @Param status query string false "Loan status"
// @Param order query string false "Order by"
// @Success 200 {object} []domain.Loan
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/loans [get]
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

// ApproveRejectLoan godoc
// @Summary Approve or reject a loan
// @Description Approve or reject a loan
// @Tags Admin
// @Accept json
// @Produce json
// @Param loan_id path string true "Loan ID"
// @Param status body string true "Loan status (approved/rejected)"
// @Success 200 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/loans/{loan_id}/status [patch]
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

// DeleteLoan godoc
// @Summary Delete a loan
// @Description Delete a loan
// @Tags Admin
// @Accept json
// @Produce json
// @Param loan_id path string true "Loan ID"
// @Success 200 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/loans/{loan_id} [delete]
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

// ViewLogs godoc
// @Summary View logs
// @Description View logs
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Log
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/logs [get]
func (lc *LoanController) ViewLogs(c *gin.Context) {
	logs, err := lc.LoanUsecase.ViewLogs(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"logs": logs})
}
