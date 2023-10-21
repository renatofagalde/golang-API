package repository

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	databaseName := "user_database_test"
	colleactionName := "user_collection_test"

	os.Setenv("MONGO_DB_COLLECTION", colleactionName)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDB.Close()

	mtestDB.Run("when_sending_a_valid_userID_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		err := repo.DeleteUser("id-mock")

		assert.Nil(t, err)
	})

	mtestDB.Run("when_sending_a_invalid_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		err := repo.DeleteUser("test-mock")

		assert.NotNil(t, err)
	})

}
