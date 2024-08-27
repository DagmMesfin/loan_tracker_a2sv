package repository

import (
	"context"
	"errors"
	"fmt"
	"loan_tracker_api/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const perpage = 10

// LoanRepository represents the loan repository contract
type LoanRepository struct {
	client *mongo.Client
	loanDB *mongo.Collection
	logDB  *mongo.Collection
}

// NewLoanRepository creates a new instance of LoanRepository
func NewLoanRepository(client *mongo.Client) domain.LoanRepository {
	return &LoanRepository{
		client: client,
		loanDB: client.Database("Loan-Tracker").Collection("Loans"),
		logDB:  client.Database("Loan-Tracker").Collection("Logs"),
	}
}

// ApplyForLoan applies for a loan
func (lr *LoanRepository) ApplyForLoan(loan *domain.Loan, userid string) error {
	useridobj, _ := primitive.ObjectIDFromHex(userid)
	loan.CreatedAt = time.Now()
	loan.UpdatedAt = time.Now()
	loan.UserID = useridobj
	loan.Status = "pending"
	loan.Interest = 0.05
	loan.ID = primitive.NewObjectID()

	log := domain.Log{
		ID:        primitive.NewObjectID(),
		UserID:    loan.UserID,
		Activity:  "Applied for a loan",
		CreatedAt: time.Now(),
	}

	_, err := lr.loanDB.InsertOne(context.Background(), loan)

	if err == nil {
		_, err = lr.logDB.InsertOne(context.Background(), log)
	}

	return err
}

// LoanDetails returns the details of a loan
func (lr *LoanRepository) LoanDetails(loanID string, userid string) (domain.Loan, error) {
	var loan domain.Loan

	loanIDObj, _ := primitive.ObjectIDFromHex(loanID)
	userIDObj, _ := primitive.ObjectIDFromHex(userid)

	err := lr.loanDB.FindOne(context.Background(), bson.M{"_id": loanIDObj, "user_id": userIDObj}).Decode(&loan)

	return loan, err
}

// ViewAllLoans returns all loans
func (lr *LoanRepository) ViewAllLoans(pgnum int, status, order string) ([]domain.Loan, int, error) {

	if pgnum == 0 {
		pgnum = 1
	}

	sorto := -1
	skip := perpage * (pgnum - 1)
	filter := bson.M{}

	if status == "" {
		status = "all"
	}

	if order == "" && status == "pending" {
		order = "asc"
	} else {
		order = "desc"
	}

	if status == "all" {
		status = ""
	} else if status == "pending" || status == "approved" || status == "rejected" {
		filter["status"] = status
	} else {
		return nil, 0, errors.New("Invalid status parameter")
	}

	if order != "asc" && order != "desc" {
		return nil, 0, errors.New("Invalid order parameter")
	}

	if order == "asc" {
		sorto = 1
	}

	count, err := lr.loanDB.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, 0, errors.New("Error counting documents")
	}

	findoptions := options.Find()
	findoptions.SetSkip(int64(skip))
	findoptions.SetLimit(perpage)
	findoptions.SetSort(bson.D{{Key: "created_at", Value: sorto}})

	var loans []domain.Loan
	cursor, err := lr.loanDB.Find(context.Background(), filter, findoptions)
	if err != nil {
		return nil, 0, errors.New("Error fetching loans")
	}

	defer cursor.Close(context.Background())
	err = cursor.All(context.Background(), &loans)

	return loans, int(count), err
}

// ApproveRejectLoan approves or rejects a loan
func (lr *LoanRepository) ApproveRejectLoan(loanID string, status string, userid string) error {
	userIDObj, _ := primitive.ObjectIDFromHex(userid)
	//findout if the loan was accepted or rejected beforehand
	var loan domain.Loan
	loanIDObj, _ := primitive.ObjectIDFromHex(loanID)
	err := lr.loanDB.FindOne(context.Background(), bson.M{"_id": loanIDObj}).Decode(&loan)
	if err != nil {
		return errors.New("Loan not found")
	}

	if loan.Status == "approved" || loan.Status == "rejected" {
		return errors.New("Loan already processed")
	}

	fmt.Println("loan", loan)

	log := domain.Log{
		ID:        primitive.NewObjectID(),
		UserID:    userIDObj,
		Activity:  "Loan " + status + "by Admin",
		CreatedAt: time.Now(),
	}

	//update the status of the loan
	res, erro := lr.loanDB.UpdateOne(context.Background(), bson.M{"_id": loanIDObj}, bson.M{"$set": bson.M{"status": status}})

	if erro == nil && res.ModifiedCount != 0 {
		_, erro = lr.logDB.InsertOne(context.Background(), log)
	}

	return erro
}

// DeleteLoan deletes a loan
func (lr *LoanRepository) DeleteLoan(loanID string, userid string) error {
	userIDObj, _ := primitive.ObjectIDFromHex(userid)
	loanIDObj, _ := primitive.ObjectIDFromHex(loanID)

	log := domain.Log{
		ID:        primitive.NewObjectID(),
		UserID:    userIDObj,
		Activity:  "Deleted a loan",
		CreatedAt: time.Now(),
	}

	res, err := lr.loanDB.DeleteOne(context.Background(), bson.M{"_id": loanIDObj})

	if err == nil && res.DeletedCount != 0 {
		fmt.Println("log", log)
		_, err = lr.logDB.InsertOne(context.Background(), log)
	}

	return err
}

// ViewLogs returns all logs
func (lr *LoanRepository) ViewLogs() ([]domain.Log, error) {
	var logs []domain.Log
	cursor, err := lr.logDB.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	err = cursor.All(context.Background(), &logs)
	return logs, err
}
