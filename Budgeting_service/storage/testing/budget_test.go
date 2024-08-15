package test

import (
	"errors"
	"testing"

	pb "budgeting/genprotos"
	"budgeting/storage/mongo"
	m "budgeting/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateBudget(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("CreateBudget_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewBudget(mt.DB)

		req := &pb.CreateBudgetRequest{
			Budget: &pb.Budget{
				Id:         primitive.NewObjectID().Hex(),
				UserId:     "user123",
				CategoryId: "category123",
				Period:     "monthly",
				Amount:     100.0,
				StartDate:  "2024-01-01",
				EndDate:    "2024-01-31",
			},
		}

		resp, err := budgetService.CreateBudget(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.CreateBudgetResponse{}, resp)
	})
}

func TestUpdateBudget(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("UpdateBudget_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewBudget(mt.DB)

		req := &pb.UpdateBudgetRequest{
			Budget: &pb.Budget{
				Id:         primitive.NewObjectID().Hex(),
				UserId:     "user123",
				CategoryId: "category123",
				Period:     "monthly",
				Amount:     150.0,
				StartDate:  "2024-01-01",
				EndDate:    "2024-01-31",
			},
		}

		resp, err := budgetService.UpdateBudget(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.UpdateBudgetResponse{}, resp)
	})
}

func TestGetBudget(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("GetBudget_Success", func(mt *mtest.T) {
		expectedBudget := m.Budget{
			ID:         "example-id",
			UserID:     "user123",
			CategoryID: "category123",
			Period:     "monthly",
			Amount:     100.0,
			StartDate:  "2024-01-01",
			EndDate:    "2024-01-31",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.budget", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedBudget.ID},
			{Key: "user_id", Value: expectedBudget.UserID},
			{Key: "category_id", Value: expectedBudget.CategoryID},
			{Key: "period", Value: expectedBudget.Period},
			{Key: "amount", Value: expectedBudget.Amount},
			{Key: "start_date", Value: expectedBudget.StartDate},
			{Key: "end_date", Value: expectedBudget.EndDate},
		}))

		budgetService := mongo.NewBudget(mt.DB)

		req := &pb.GetBudgetRequest{Id: expectedBudget.ID}

		resp, err := budgetService.GetBudget(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.GetBudgetResponse{
			Budget: &pb.Budget{
				Id:         expectedBudget.ID,
				UserId:     expectedBudget.UserID,
				CategoryId: expectedBudget.CategoryID,
				Period:     expectedBudget.Period,
				Amount:     expectedBudget.Amount,
				StartDate:  expectedBudget.StartDate,
				EndDate:    expectedBudget.EndDate,
			},
		}, resp)
	})

	mt.Run("GetBudget_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "test.budget", mtest.FirstBatch))

		budgetService := mongo.NewBudget(mt.DB)

		req := &pb.GetBudgetRequest{Id: "nonexistent-id"}

		resp, err := budgetService.GetBudget(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, errors.New("budget not found"), err)
	})

	mt.Run("GetBudget_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		budgetService := mongo.NewBudget(mt.DB)

		req := &pb.GetBudgetRequest{Id: "example-id"}

		resp, err := budgetService.GetBudget(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}


func TestDeleteBudget(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("DeleteBudget_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewBudget(mt.DB)

		req := &pb.DeleteBudgetRequest{Id: "example-id"}

		resp, err := budgetService.DeleteBudget(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteBudgetResponse{}, resp)
	})

	mt.Run("DeleteBudget_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewBudget(mt.DB)

		req := &pb.DeleteBudgetRequest{Id: "nonexistent-id"}

		resp, err := budgetService.DeleteBudget(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteBudgetResponse{}, resp)
	})

	mt.Run("DeleteBudget_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		budgetService := mongo.NewBudget(mt.DB)

		req := &pb.DeleteBudgetRequest{Id: "example-id"}

		resp, err := budgetService.DeleteBudget(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "internal server error")
	})
}
