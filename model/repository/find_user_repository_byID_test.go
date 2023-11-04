package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang-API/model/repository/entity"
	"golang-API/model/repository/entity/convert"
	"os"
	"testing"
)

func TestUserRepository_FindUserByID(t *testing.T) {
	databaseName := "user_database_test"
	colleactionName := "user_collection_test"

	os.Setenv("MONGO_DB_COLLECTION", colleactionName)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDB.Close()

	mtestDB.Run("when_sending_a_valid_ID_returns_success", func(mt *mtest.T) {

		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    EMAIL,
			Password: "teste",
			Name:     "teste",
			Age:      18,
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, colleactionName), mtest.FirstBatch, convert.ConvertEntityToBson(userEntity)))
		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.NotNil(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.NotNil(t, userDomain.GetEmail(), userEntity.Email)
		assert.NotNil(t, userDomain.GetName(), userEntity.Name)
		assert.NotNil(t, userDomain.GetPassword(), userEntity.Password)

	})

	mtestDB.Run("when_sending_a_invalid_ID_returns_error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		userDomain, err := repo.FindUserByID("_erro")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDB.Run("when_sending_an_ID_absent", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, colleactionName), mtest.FirstBatch))

		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		userDomain, err := repo.FindUserByID("absent")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
