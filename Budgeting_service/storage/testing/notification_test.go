package test

import (
	"testing"

	pb "budgeting/genprotos"
	"budgeting/storage/mongo"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateNotification(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("CreateNotification_Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		budgetService := mongo.NewNotification(mt.DB)

		req := &pb.CreateNotificationRequest{
			Notification: &pb.Notification{
				Id:      primitive.NewObjectID().Hex(),
				UserId:  "user123",
				Message: "nma gap",
			},
		}

		resp, err := budgetService.CreateNotification(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, &pb.CreateNotificationResponse{}, resp)
	})

}
