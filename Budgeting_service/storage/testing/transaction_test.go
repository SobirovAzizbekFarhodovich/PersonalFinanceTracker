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

func TestCreateTransaction(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("CreateTransaction_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewTransaction(mt.DB)

		req := &pb.CreateTransactionRequest{
			Transaction: &pb.Transaction{
				Id:          primitive.NewObjectID().Hex(),
				UserId:      "user123",
				CategoryId:  "category123",
				AccountId:   "monthly",
				Amount:      100.0,
				Type:        "income",
				Description: "farqi yo",
				Date:        "2024-01-31",
			},
		}

		resp, err := budgetService.CreateTransaction(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.CreateTransactionResponse{}, resp)
	})

}

func TestUpdateTransaction(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("UpdateTransaction_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewTransaction(mt.DB)

		req := &pb.UpdateTransactionRequest{
			Transaction: &pb.Transaction{
				Id:          primitive.NewObjectID().Hex(),
				UserId:      "user123",
				CategoryId:  "category123",
				AccountId:   "monthly",
				Amount:      100.0,
				Type:        "income",
				Description: "farqi yo",
				Date:        "2024-01-31",
			},
		}

		resp, err := budgetService.UpdateTransaction(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.UpdateTransactionResponse{}, resp)
	})
}

func TestGetTransaction(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("GetTransaction_Success", func(mt *mtest.T) {
		expectedTransaction := m.Transaction{
			ID:          "example-id",
			UserID:      "user123",
			CategoryID:  "category123",
			AccountID:   "1",
			Amount:      100.0,
			Type:        "income",
			Date:        "2024-01-31",
			Description: "farqi yo",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.transaction", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedTransaction.ID},
			{Key: "user_id", Value: expectedTransaction.UserID},
			{Key: "category_id", Value: expectedTransaction.CategoryID},
			{Key: "account_id", Value: expectedTransaction.AccountID},
			{Key: "amount", Value: expectedTransaction.Amount},
			{Key: "type", Value: expectedTransaction.Type},
			{Key: "date", Value: expectedTransaction.Date},
			{Key: "description", Value: expectedTransaction.Description},
		}))

		transactionService := mongo.NewTransaction(mt.DB)

		req := &pb.GetTransactionRequest{Id: expectedTransaction.ID}

		resp, err := transactionService.GetTransaction(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.GetTransactionResponse{
			Transaction: &pb.Transaction{
				Id:          expectedTransaction.ID,
				UserId:      expectedTransaction.UserID,
				CategoryId:  expectedTransaction.CategoryID,
				AccountId:   expectedTransaction.AccountID,
				Amount:      expectedTransaction.Amount,
				Type:        expectedTransaction.Type,
				Description: expectedTransaction.Description,
				Date:        expectedTransaction.Date,
			},
		}, resp)
	})

	// The rest of your test cases remain unchanged
	mt.Run("GetTransaction_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "test.transaction", mtest.FirstBatch))

		transactionService := mongo.NewTransaction(mt.DB)

		req := &pb.GetTransactionRequest{Id: "nonexistent-id"}

		resp, err := transactionService.GetTransaction(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, errors.New("Transaction not found"), err)
	})

	mt.Run("GetTransaction_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		transactionService := mongo.NewTransaction(mt.DB)

		req := &pb.GetTransactionRequest{Id: "example-id"}

		resp, err := transactionService.GetTransaction(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}


func TestDeleteTransaction(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("DeleteTransaction_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		TransactionService := mongo.NewTransaction(mt.DB)

		req := &pb.DeleteTransactionRequest{Id: "example-id"}

		resp, err := TransactionService.DeleteTransaction(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteTransactionResponse{}, resp)
	})

	mt.Run("DeleteTransaction_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		TransactionService := mongo.NewTransaction(mt.DB)

		req := &pb.DeleteTransactionRequest{Id: "nonexistent-id"}

		resp, err := TransactionService.DeleteTransaction(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteTransactionResponse{}, resp)
	})

	mt.Run("DeleteTransaction_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		TransactionService := mongo.NewTransaction(mt.DB)

		req := &pb.DeleteTransactionRequest{Id: "example-id"}

		resp, err := TransactionService.DeleteTransaction(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "internal server error")
	})
}
