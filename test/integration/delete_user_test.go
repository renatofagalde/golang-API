package tests

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestDeleteUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
	ID := primitive.NewObjectID()

	_, err := Database.
		Collection(os.Getenv("MONGO_DB_COLLECTION")).
		InsertOne(context.Background(), bson.M{"_id": ID, "name": t.Name(), "email": "test@test.com"})

	if err != nil {
		t.Fatal(err)
		return
	}

	params := []gin.Param{
		{
			Key:   "id",
			Value: ID.Hex(),
		},
	}

	MakeRequest(ctx, "DELETE", url.Values{}, params, nil)
	UserController.Delete(ctx)

	assert.EqualValues(t, http.StatusOK, recorder.Code)

	filter := bson.D{{Key: "_id", Value: ID}}
	result := Database.
		Collection(os.Getenv("MONGO_DB_COLLECTION")).
		FindOne(context.Background(), filter)

	assert.NotNil(t, result.Err())

}
