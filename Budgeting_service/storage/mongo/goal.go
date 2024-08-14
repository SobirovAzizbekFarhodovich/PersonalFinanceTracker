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
	"go.mongodb.org/mongo-driver/mongo/options"
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
		return nil, errors.New("goal not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.GetGoalResponse{Goal: e.BsonToGoal(&bsonGoal)}, nil
}

func (s *Goal) ListGoals(req *pb.ListGoalsRequest) (*pb.ListGoalsResponse, error) {
	limit := req.Limit
	page := req.Page
	skip := (page - 1) * limit

	cursor, err := s.mongo.Find(context.TODO(), bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var accounts []*pb.Goal
	for cursor.Next(context.TODO()) {
		var bsonAccount m.Goal
		if err := cursor.Decode(&bsonAccount); err != nil {
			return nil, err
		}
		account := e.BsonToGoal(&bsonAccount)
		accounts = append(accounts, account)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListGoalsResponse{
		Goals: accounts,
	}, nil
}

func (s *Goal) GenerateGoalProgressReport(req *pb.GenerateGoalProgressReportRequest) (*pb.GenerateGoalProgressReportResponse, error) {
	var bsonGoal m.Goal
	err := s.mongo.FindOne(context.TODO(), bson.M{"_id": req.Id}).Decode(&bsonGoal)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("goal not found")
	} else if err != nil {
		return nil, err
	}
	goal := e.BsonToGoal(&bsonGoal)

	deadlineStr := goal.Deadline + "T00:00:00Z"
	deadlineTime, err := time.Parse(time.RFC3339, deadlineStr)
	if err != nil {
		return nil, fmt.Errorf("invalid deadline format")
	}

	remainAmount := goal.TargetAmount - goal.CurrentAmount
	if remainAmount < 0 {
		remainAmount = 0
	}

	status := "in_progress"
	
	if time.Now().After(deadlineTime) {
		status = "failed"
	}

	resp := &pb.GenerateGoalProgressReportResponse{
		UserId:        goal.UserId,
		Name:          goal.Name,
		TargetAmount:  goal.TargetAmount,
		CurrentAmount: goal.CurrentAmount,
		RemainAmount:  remainAmount,
		Deadline:      goal.Deadline,
		Status:        status,
	}

	return resp, nil
}

