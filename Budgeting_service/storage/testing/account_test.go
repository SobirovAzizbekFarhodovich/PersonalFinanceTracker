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

func TestCreateAccount(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("CreateAccount_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewAccount(mt.DB)

		req := &pb.CreateAccountRequest{
			Account: &pb.Account{
				Id:       primitive.NewObjectID().Hex(),
				UserId:   "user123",
				Name:     "category123",
				Currency: "monthly",
				Balance:  100.0,
				Type:     "income",
			},
		}

		resp, err := budgetService.CreateAccount(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.CreateAccountResponse{}, resp)
	})

}

func TestUpdateAccount(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("UpdateAccount_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewAccount(mt.DB)

		req := &pb.UpdateAccountRequest{
			Account: &pb.Account{
				Id:       primitive.NewObjectID().Hex(),
				UserId:   "user123",
				Name:     "category123",
				Currency: "monthly",
				Balance:  100.0,
				Type:     "income",
			},
		}

		resp, err := budgetService.UpdateAccount(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.UpdateAccountResponse{}, resp)
	})
}

func TestGetAccount(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("GetAccount_Success", func(mt *mtest.T) {
		expectedAccount := m.Account{
			ID:       "example-id",
			UserID:   "user123",
			Name:     "category123",
			Type:     "monthly",
			Balance:  100.0,
			Currency: "2024-01-01",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.account", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedAccount.ID},
			{Key: "user_id", Value: expectedAccount.UserID},
			{Key: "name", Value: expectedAccount.Name},
			{Key: "type", Value: expectedAccount.Type},
			{Key: "balance", Value: expectedAccount.Balance},
			{Key: "currency", Value: expectedAccount.Currency},
		}))

		accountService := mongo.NewAccount(mt.DB)

		req := &pb.GetAccountRequest{Id: expectedAccount.ID}

		resp, err := accountService.GetAccount(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.GetAccountResponse{
			Account: &pb.Account{
				Id:       expectedAccount.ID,
				UserId:   expectedAccount.UserID,
				Name:     expectedAccount.Name,
				Type:     expectedAccount.Type,
				Balance:  expectedAccount.Balance,
				Currency: expectedAccount.Currency,
			},
		}, resp)
	})

	mt.Run("GetAccount_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "test.account", mtest.FirstBatch))

		accountService := mongo.NewAccount(mt.DB)

		req := &pb.GetAccountRequest{Id: "nonexistent-id"}

		resp, err := accountService.GetAccount(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, errors.New("Account not found"), err)
	})

	mt.Run("GetAccount_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		accountService := mongo.NewAccount(mt.DB)

		req := &pb.GetAccountRequest{Id: "example-id"}

		resp, err := accountService.GetAccount(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestDeleteAccount(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("DeleteAccount_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewAccount(mt.DB)

		req := &pb.DeleteAccountRequest{Id: "example-id"}

		resp, err := budgetService.DeleteAccount(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteAccountResponse{}, resp)
	})

	mt.Run("DeleteAccount_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewAccount(mt.DB)

		req := &pb.DeleteAccountRequest{Id: "nonexistent-id"}

		resp, err := budgetService.DeleteAccount(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteAccountResponse{}, resp)
	})

	mt.Run("DeleteAccount_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		budgetService := mongo.NewAccount(mt.DB)

		req := &pb.DeleteAccountRequest{Id: "example-id"}

		resp, err := budgetService.DeleteAccount(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "internal server error")
	})
}
