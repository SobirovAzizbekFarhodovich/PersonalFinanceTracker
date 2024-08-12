package service

import (
	pb "budgeting/genprotos"
	"budgeting/storage"
	"context"

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
	return &pb.CreateGoalResponse{}, nil
}

func (s *GoalService) UpdateGoal(c context.Context, req *pb.UpdateGoalRequest) (*pb.UpdateGoalResponse, error) {
	_, err := s.stg.Goal().UpdateGoal(req)
	if err != nil {
		return nil, err
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
	return res, nil
}

func (s *GoalService) ListGoals(c context.Context, req *pb.ListGoalsRequest) (*pb.ListGoalsResponse, error) {
	res, err := s.stg.Goal().ListGoals(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *GoalService) GenerateGoalProgressReport(c context.Context, req *pb.GenerateGoalProgressReportRequest) (*pb.GenerateGoalProgressReportResponse, error) {
	res, err := s.stg.Goal().GenerateGoalProgressReport(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
