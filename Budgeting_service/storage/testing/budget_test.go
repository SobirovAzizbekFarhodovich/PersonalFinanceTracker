package test

import (
	"testing"

	pb "budgeting/genprotos"
	"budgeting/storage/mongo"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateBudget(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	// Setup the mock collection
	mt.Run("CreateBudget_Success", func(mt *mtest.T) {
		// Mock insert result
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		// Create a new Budget instance with the mock collection
		budgetService := mongo.NewBudget(mt.DB)

		// Create a mock request
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

		// Call the CreateBudget method
		resp, err := budgetService.CreateBudget(req)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.CreateBudgetResponse{}, resp)
	})

	mt.Run("CreateBudget_InsertFail", func(mt *mtest.T) {
		// Mock insert error
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "duplicate key error",
		}))

		// Create a new Budget instance with the mock collection
		budgetService := mongo.NewBudget(mt.DB)

		// Create a mock request
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

		// Call the CreateBudget method
		resp, err := budgetService.CreateBudget(req)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
