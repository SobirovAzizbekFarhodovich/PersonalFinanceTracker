package service

import (
	pb "budgeting/genprotos"
	"budgeting/storage"
	"context"

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
	_, err := s.stg.Budget().CreateBudget(req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBudgetResponse{}, nil
}

func (s *BudgetService) UpdateBudget(c context.Context, req *pb.UpdateBudgetRequest) (*pb.UpdateBudgetResponse, error) {
	_, err := s.stg.Budget().UpdateBudget(req)
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
	return res, nil
}
