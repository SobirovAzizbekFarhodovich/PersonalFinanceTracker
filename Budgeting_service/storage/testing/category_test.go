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

func TestCreateCategory(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("CreateCategory_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewCategory(mt.DB)

		req := &pb.CreateCategoryRequest{
			Category: &pb.Category{
				Id:     primitive.NewObjectID().Hex(),
				UserId: "user123",
				Name:   "category123",
				Type:   "income",
			},
		}

		resp, err := budgetService.CreateCategory(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.CreateCategoryResponse{}, resp)
	})

}

func TestUpdateCategory(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("UpdateCategory_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewCategory(mt.DB)

		req := &pb.UpdateCategoryRequest{
			Category: &pb.Category{
				Id:     primitive.NewObjectID().Hex(),
				UserId: "user123",
				Name:   "category123",
				Type:   "income",
			},
		}

		resp, err := budgetService.UpdateCategory(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.UpdateCategoryResponse{}, resp)
	})
}

func TestGetCategory(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("GetCategory_Success", func(mt *mtest.T) {
		expectedAccount := m.Category{
			ID:       "example-id",
			UserID:   "user123",
			Name:     "category123",
			Type:     "monthly",
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.category", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedAccount.ID},
			{Key: "user_id", Value: expectedAccount.UserID},
			{Key: "name", Value: expectedAccount.Name},
			{Key: "type", Value: expectedAccount.Type},
		}))

		accountService := mongo.NewCategory(mt.DB)

		req := &pb.GetCategoryRequest{Id: expectedAccount.ID}

		resp, err := accountService.GetCategory(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.GetCategoryResponse{
			Category: &pb.Category{
				Id:       expectedAccount.ID,
				UserId:   expectedAccount.UserID,
				Name:     expectedAccount.Name,
				Type:     expectedAccount.Type,
			},
		}, resp)
	})

	mt.Run("GetCategory_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "test.category", mtest.FirstBatch))
	
		categoryService := mongo.NewCategory(mt.DB)
	
		req := &pb.GetCategoryRequest{Id: "nonexistent-id"}
	
		resp, err := categoryService.GetCategory(req)
	
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, errors.New("category not found"), err) 
	})
	
	mt.Run("GetCategory_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		accountService := mongo.NewCategory(mt.DB)

		req := &pb.GetCategoryRequest{Id: "example-id"}

		resp, err := accountService.GetCategory(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestDeleteCategory(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("DeleteCategory_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewCategory(mt.DB)

		req := &pb.DeleteCategoryRequest{Id: "example-id"}

		resp, err := budgetService.DeleteCategory(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteCategoryResponse{}, resp)
	})

	mt.Run("DeleteCategory_NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewCategory(mt.DB)

		req := &pb.DeleteCategoryRequest{Id: "nonexistent-id"}

		resp, err := budgetService.DeleteCategory(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.IsType(t, &pb.DeleteCategoryResponse{}, resp)
	})

	mt.Run("DeleteCategory_Error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "internal server error",
		}))

		budgetService := mongo.NewCategory(mt.DB)

		req := &pb.DeleteCategoryRequest{Id: "example-id"}

		resp, err := budgetService.DeleteCategory(req)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Contains(t, err.Error(), "internal server error")
	})
}
