package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	e "budgeting/extra"
	pb "budgeting/genprotos"
	m "budgeting/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Goal struct {
	mongo *mongo.Collection
}

func NewGoal(db *mongo.Database) *Goal {
	return &Goal{mongo: db.Collection("goal")}
}

func (s *Goal) CreateGoal(req *pb.CreateGoalRequest) (*pb.CreateGoalResponse, error) {
	bsonGoal := e.GoalToBSON(req.Goal)
	_, err := s.mongo.InsertOne(context.TODO(), bsonGoal)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGoalResponse{}, nil
}

func (s *Goal) UpdateGoal(req *pb.UpdateGoalRequest) (*pb.UpdateGoalResponse, error) {
	bsonGoal := e.GoalToBSON(req.Goal)
	filter := bson.M{"_id": bsonGoal.ID}
	update := bson.M{"$set": bsonGoal}

	_, err := s.mongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateGoalResponse{}, nil
}

func (s *Goal) DeleteGoal(req *pb.DeleteGoalRequest) (*pb.DeleteGoalResponse, error) {
	_, err := s.mongo.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGoalResponse{}, nil
}

func (s *Goal) GetGoal(req *pb.GetGoalRequest) (*pb.GetGoalResponse, error) {
	var bsonGoal m.Goal
	err := s.mongo.FindOne(context.TODO(), bson.M{"_id": req.Id}).Decode(&bsonGoal)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("progress not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.GetGoalResponse{Goal: e.BsonToGoal(&bsonGoal)}, nil
}

func (s *Goal) ListGoals(req *pb.ListGoalsRequest) (*pb.ListGoalsResponse, error) {
	limit := req.Limit
	page := req.Page
	skip := (page - 1) * limit

	pipeline := []bson.M{
		{"$skip": skip},
		{"$limit": limit},
		{"$project": bson.M{
			"_id":            0,
			"user_id":        1,
			"name":           1,
			"target_amount":  1,
			"current_amount": 1,
			"deadline":       1,
			"status":         1,
		}},
	}

	cursor, err := s.mongo.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var Goals []*pb.Goal
	for cursor.Next(context.TODO()) {
		var Goal pb.Goal
		if err := cursor.Decode(&Goal); err != nil {
			return nil, err
		}
		Goals = append(Goals, &Goal)
	}

	return &pb.ListGoalsResponse{
		Goals: Goals,
	}, nil
}

func (s *Goal) GenerateGoalProgressReport(req *pb.GenerateGoalProgressReportRequest) (*pb.GenerateGoalProgressReportResponse, error) {
	var goal pb.Goal
	err := s.mongo.FindOne(context.TODO(), bson.M{"_id": req.Id}).Decode(&goal)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("goal not found")
		}
		return nil, err
	}

	remainAmount := goal.TargetAmount - goal.CurrentAmount
	if remainAmount < 0 {
		remainAmount = 0
	}

	deadlineTime, err := time.Parse(time.RFC3339, goal.Deadline)
	if err != nil {
		return nil, fmt.Errorf("invalid deadline format")
	}

	status := "in_progress"
	if goal.CurrentAmount >= goal.TargetAmount {
		status = "achieved"
	} else if time.Now().After(deadlineTime) {
		status = "failed"
	}

	// Create the response
	resp := &pb.GenerateGoalProgressReportResponse{
		UserId:        goal.UserId,
		Name:          goal.Name,
		TargetAmount:  float32(goal.TargetAmount),
		CurrentAmount: float32(goal.CurrentAmount),
		RemainAmount:  remainAmount,
		Deadline:      goal.Deadline,
		Status:        status,
	}

	return resp, nil
}
