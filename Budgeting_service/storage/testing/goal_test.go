package test

import (
	"errors"
	"testing"

	pb "budgeting/genprotos"
	m "budgeting/models"
	"budgeting/storage/mongo"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateGoal(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("CreateGoal_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewGoal(mt.DB)

		req := &pb.CreateGoalRequest{
			Goal: &pb.Goal{
				Id:            primitive.NewObjectID().Hex(),
				UserId:        "user123",
				CurrentAmount: 10000,
				TargetAmount:  100000,
				Status:        "in_progress",
				Deadline:      "2024-01-31",
			},
		}

		resp, err := budgetService.CreateGoal(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.CreateGoalResponse{}, resp)
	})

}

func TestUpdateGoal(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("UpdateGoal_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewGoal(mt.DB)

		req := &pb.UpdateGoalRequest{
			Goal: &pb.Goal{
				Id:            primitive.NewObjectID().Hex(),
				UserId:        "user123",
				CurrentAmount: 10000,
				TargetAmount:  100000,
				Status:        "in_progress",
				Deadline:      "2024-01-31",
			},
		}

		resp, err := budgetService.UpdateGoal(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.UpdateGoalResponse{}, resp)
	})
}

func TestGetGoal(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("GetGoal_Success", func(mt *mtest.T) {
		expectedAccount := m.Goal{
			ID:            "example-id",
			UserID:        "user123",
			Name:          "category123",
			TargetAmount:  100000,
			CurrentAmount: 100.0,
			Deadline:      "2024-01-01",
			Status:        "nma gap",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.goal", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedAccount.ID},
			{Key: "user_id", Value: expectedAccount.UserID},
			{Key: "name", Value: expectedAccount.Name},
			{Key: "target_amount", Value: expectedAccount.TargetAmount},
			{Key: "current_amount", Value: expectedAccount.CurrentAmount},
			{Key: "deadline", Value: expectedAccount.Deadline},
			{Key: "status", Value: expectedAccount.Status},
		}))

		accountService := mongo.NewGoal(mt.DB)

		req := &pb.GetGoalRequest{Id: expectedAccount.ID}

		resp, err := accountService.GetGoal(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.GetGoalResponse{
			Goal: &pb.Goal{
				Id:            expectedAccount.ID,
				UserId:        expectedAccount.UserID,
				Name:          expectedAccount.Name,
				TargetAmount:  expectedAccount.TargetAmount,
				CurrentAmount: expectedAccount.CurrentAmount,
				Deadline:      expectedAccount.Deadline,
				Status:        expectedAccount.Status,
			},
		}, resp)
	})

	mt.Run("GetGoal_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "test.goal", mtest.FirstBatch))

		accountService := mongo.NewGoal(mt.DB)

		req := &pb.GetGoalRequest{Id: "nonexistent-id"}

		resp, err := accountService.GetGoal(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, errors.New("goal not found"), err)
	})

	mt.Run("GetGoal_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		accountService := mongo.NewGoal(mt.DB)

		req := &pb.GetGoalRequest{Id: "example-id"}

		resp, err := accountService.GetGoal(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}


func TestDeleteGoal(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("DeleteGoal_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		GoalService := mongo.NewGoal(mt.DB)

		req := &pb.DeleteGoalRequest{Id: "example-id"}

		resp, err := GoalService.DeleteGoal(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteGoalResponse{}, resp)
	})

	mt.Run("DeleteGoal_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		GoalService := mongo.NewGoal(mt.DB)

		req := &pb.DeleteGoalRequest{Id: "nonexistent-id"}

		resp, err := GoalService.DeleteGoal(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteGoalResponse{}, resp)
	})

	mt.Run("DeleteGoal_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		GoalService := mongo.NewGoal(mt.DB)

		req := &pb.DeleteGoalRequest{Id: "example-id"}

		resp, err := GoalService.DeleteGoal(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "internal server error")
	})
}
