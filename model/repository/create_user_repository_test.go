package repository

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang-basic/model"
	"os"
	"testing"
)

const (
	EMAIL = "teste@teste.com"
)

func TestUserRepository_CreateUser(t *testing.T) {
	databaseName := "user_database_test"
	colleactionName := "user_collection_test"

	os.Setenv("MONGO_DB_COLLECTION", colleactionName)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDB.Close()

	mtestDB.Run("when_sending_a_valid_domain_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		domain := model.NewUserDomain(EMAIL, "password", "test", 18)
		userDomain, err := repo.CreateUser(domain)

		_, errID := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errID) //validando o ID de forma indireta
		assert.NotNil(t, userDomain.GetEmail(), domain.GetEmail())
		assert.NotNil(t, userDomain.GetName(), domain.GetName())
		assert.NotNil(t, userDomain.GetPassword(), domain.GetPassword())

	})

	mtestDB.Run("when_sending_a_invalid_domain_return_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		domain := model.NewUserDomain(EMAIL, "password", "test", 18)
		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain) //validando o ID de forma indireta

	})

}
