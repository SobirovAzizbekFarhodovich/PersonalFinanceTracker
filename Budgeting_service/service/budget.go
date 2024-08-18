package service

import (
	pb "budgeting/genprotos"
	"budgeting/storage"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BudgetService struct {
	stg storage.StorageI
	pb.UnimplementedBudgetServiceServer
}

func NewBudgetService(stg *storage.StorageI) *BudgetService {
	return &BudgetService{stg: *stg}
}

func (s *BudgetService) CreateBudget(c context.Context, req *pb.CreateBudgetRequest) (*pb.CreateBudgetResponse, error) {
	id := uuid.NewString()
	req.Budget.Id = id

	if _, err := uuid.Parse(req.Budget.UserId); err != nil {
		return nil, fmt.Errorf("invalid UserId: must be a valid UUID")
	}

	if _, err := uuid.Parse(req.Budget.CategoryId); err != nil {
		return nil, fmt.Errorf("invalid CategoryId: must be a valid UUID")
	}

	if req.Budget.Period != "daily" && req.Budget.Period != "weekly" && req.Budget.Period != "monthly" && req.Budget.Period != "yearly" {
		return nil, fmt.Errorf("invalid Budget Period: must be 'daily' or 'weekly' or 'monthly' or 'yearly'")
	}

	_, err := time.Parse("2006-01-02", req.Budget.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid EndDate: must be a valid date in 'YYYY-MM-DD' format")
	}

	_, err = time.Parse("2006-01-02", req.Budget.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid StartDate: must be a valid date in 'YYYY-MM--DD' format")
	}

	_, err = s.stg.Budget().CreateBudget(req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBudgetResponse{}, nil
}

func (s *BudgetService) UpdateBudget(c context.Context, req *pb.UpdateBudgetRequest) (*pb.UpdateBudgetResponse, error) {
	if _, err := uuid.Parse(req.Budget.UserId); err != nil {
		return nil, fmt.Errorf("invalid UserId: must be a valid UUID")
	}

	if _, err := uuid.Parse(req.Budget.CategoryId); err != nil {
		return nil, fmt.Errorf("invalid CategoryId: must be a valid UUID")
	}

	if req.Budget.Period != "daily" && req.Budget.Period != "weekly" && req.Budget.Period != "monthly" && req.Budget.Period != "yearly" {
		return nil, fmt.Errorf("invalid Budget Period: must be 'daily' or 'weekly' or 'monthly' or 'yearly'")
	}

	_, err := time.Parse("2006-01-02", req.Budget.EndDate)
	if err != nil {
		return nil, fmt.Errorf("invalid EndDate: must be a valid date in 'YYYY-MM-DD' format")
	}

	_, err = time.Parse("2006-01-02", req.Budget.StartDate)
	if err != nil {
		return nil, fmt.Errorf("invalid StartDate: must be a valid date in 'YYYY-MM--DD' format")
	}

	_, err = s.stg.Budget().UpdateBudget(req)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateBudgetResponse{}, nil
}

func (s *BudgetService) DeleteBudget(c context.Context, req *pb.DeleteBudgetRequest) (*pb.DeleteBudgetResponse, error) {
	_, err := s.stg.Budget().DeleteBudget(req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBudgetResponse{}, nil
}

func (s *BudgetService) GetBudget(c context.Context, req *pb.GetBudgetRequest) (*pb.GetBudgetResponse, error) {
	res, err := s.stg.Budget().GetBudget(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *BudgetService) ListBudgets(c context.Context, req *pb.ListBudgetsRequest) (*pb.ListBudgetsResponse, error) {
	res, err := s.stg.Budget().ListBudgets(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *BudgetService) GenerateBudgetPerformanceReport(c context.Context, req *pb.GenerateBudgetPerformanceReportRequest) (*pb.GenerateBudgetPerformanceReportResponse, error) {
	res, err := s.stg.Budget().GenerateBudgetPerformanceReport(req)
	if err != nil {
		return nil, err
	}

	transactionres, err := s.stg.Transaction().ListTransactions(&pb.ListTransactionsRequest{})
	if err != nil {
		return nil, errors.New("failed to retrieve transactions")
	}

	var spentAmount float32 = 0
	for _, transaction := range transactionres.Transactions {
		if transaction.Type == "expense" {
			spentAmount += transaction.Amount
		}
	}

	resp := &pb.GenerateBudgetPerformanceReportResponse{
		Id:          res.Id,
		UserId:      res.UserId,
		CategoryId:  res.CategoryId,
		Amount:      res.Amount,
		Period:      res.Period,
		StartDate:   res.StartDate,
		EndDate:     res.EndDate,
		SpentAmount: spentAmount,
	}

	if spentAmount >= res.Amount {
		message := fmt.Sprintf("You have exceeded your budget by spending %.2f sum.", spentAmount)
		notificationReq := &pb.CreateNotificationRequest{
			Notification: &pb.Notification{
				Id:      uuid.NewString(),
				UserId:  res.UserId,
				Message: message,
			},
		}
		_, err = s.stg.Notification().CreateNotification(&pb.CreateNotificationRequest{Notification: notificationReq.Notification})
		if err != nil {
			return nil, errors.New("notification not created")
		}
	}

	return resp, nil
}
