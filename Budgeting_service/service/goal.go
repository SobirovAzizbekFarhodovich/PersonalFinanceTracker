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

type GoalService struct {
	stg storage.StorageI
	pb.UnimplementedGoalServiceServer
}

func NewGoalService(stg *storage.StorageI) *GoalService {
	return &GoalService{stg: *stg}
}

func (s *GoalService) CreateGoal(c context.Context, req *pb.CreateGoalRequest) (*pb.CreateGoalResponse, error) {
	id := uuid.NewString()
	req.Goal.Id = id
	_, err := s.stg.Goal().CreateGoal(req)
	if err != nil {
		return nil, err
	}

	user_id, err := s.stg.Account().GetAmount(&pb.GetAmountRequest{UserId: req.Goal.UserId})
	if err != nil {
		return nil, errors.New("user not found")
	}
	if req.Goal.TargetAmount <= user_id.Balance {
		return nil, errors.New("current amount greater than target amount")
	}
	req.Goal.Status = "in_progress"

	req.Goal.CurrentAmount = user_id.Balance

	_, err = s.stg.Goal().CreateGoal(&pb.CreateGoalRequest{Goal: req.Goal})
	if err != nil {
		return nil, errors.New("failed to create goal")
	}

	return &pb.CreateGoalResponse{}, nil
}

func (s *GoalService) UpdateGoal(c context.Context, req *pb.UpdateGoalRequest) (*pb.UpdateGoalResponse, error) {
	_, err := s.stg.Goal().UpdateGoal(req)
	if err != nil {
		return nil, err
	}

	user_id, err := s.stg.Account().GetAmount(&pb.GetAmountRequest{UserId: req.Goal.UserId})
	if err != nil {
		return nil, errors.New("user not found")
	}

	if req.Goal.TargetAmount <= user_id.Balance {
		return nil, errors.New("current amount greater than target amount")
	}
	req.Goal.Status = "in_progress"

	req.Goal.CurrentAmount = user_id.Balance
	if req.Goal.CurrentAmount >= req.Goal.TargetAmount {
		req.Goal.Status = "achieved"
	}
	deadlineStr := req.Goal.Deadline + "T00:00:00Z"
	deadlineTime, err := time.Parse(time.RFC3339, deadlineStr)
	if err != nil {
		return nil, fmt.Errorf("invalid deadline format")
	}
	if time.Now().After(deadlineTime) {
		req.Goal.Status = "failed"
	}

	return &pb.UpdateGoalResponse{}, nil
}

func (s *GoalService) DeleteGoal(c context.Context, req *pb.DeleteGoalRequest) (*pb.DeleteGoalResponse, error) {
	_, err := s.stg.Goal().DeleteGoal(req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGoalResponse{}, nil
}

func (s *GoalService) GetGoal(c context.Context, req *pb.GetGoalRequest) (*pb.GetGoalResponse, error) {
	res, err := s.stg.Goal().GetGoal(req)
	if err != nil {
		return nil, err
	}

	accountres, err := s.stg.Account().GetAmount(&pb.GetAmountRequest{UserId: res.Goal.UserId})
	if err != nil {
		return nil, errors.New("account not found")
	}

	res.Goal.CurrentAmount = accountres.Balance

	return res, nil
}

func (s *GoalService) ListGoals(c context.Context, req *pb.ListGoalsRequest) (*pb.ListGoalsResponse, error) {
	res, err := s.stg.Goal().ListGoals(req)
	if err != nil {
		return nil, err
	}

	for _, goal := range res.Goals {
		accountres, err := s.stg.Account().GetAmount(&pb.GetAmountRequest{UserId: goal.UserId})
		if err != nil {
			return nil, errors.New("account not found")
		}

		goal.CurrentAmount = accountres.Balance
	}
	
	return res, nil
}

func (s *GoalService) GenerateGoalProgressReport(c context.Context, req *pb.GenerateGoalProgressReportRequest) (*pb.GenerateGoalProgressReportResponse, error) {
	res, err := s.stg.Goal().GenerateGoalProgressReport(req)
	if err != nil {
		return nil, err
	}

	accountres, err := s.stg.Account().GetAmount(&pb.GetAmountRequest{UserId: res.UserId})
	if err != nil {
		return nil, errors.New("account not found")
	}

	res.CurrentAmount = accountres.Balance

	res.CurrentAmount = accountres.Balance
    res.RemainAmount = res.TargetAmount - res.CurrentAmount
    if res.CurrentAmount >= res.TargetAmount{
        res.Status = "achieved"
    }

	return res, nil
}
