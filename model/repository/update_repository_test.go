package repository

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang-basic/model"
	"os"
	"testing"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	colleactionName := "user_collection_test"

	os.Setenv("MONGO_DB_COLLECTION", colleactionName)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDB.Close()

	mtestDB.Run("when_sending_a_valid_user_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		domain := model.NewUserDomain(EMAIL, "password", "test", 18)
		err := repo.UpdateUser(domain.GetID(), domain)

		assert.Nil(t, err)
	})

	mtestDB.Run("when_sending_a_invalid_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		domain := model.NewUserDomain(EMAIL, "password", "test", 18)
		err := repo.UpdateUser(domain.GetID(), domain)

		assert.NotNil(t, err)
	})

}
