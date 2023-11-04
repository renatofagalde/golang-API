package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang-API/config/rest_err"
	"golang-API/model"
	"golang-API/model/repository/entity"
	"golang-API/model/repository/entity/convert"
	"os"
	"testing"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	colleactionName := "user_collection_test"

	os.Setenv("MONGO_DB_COLLECTION", colleactionName)
	defer os.Clearenv()

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mtestDB.Close()

	mtestDB.Run("when_sending_a_valid_email_returns_success", func(mt *mtest.T) {

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
		userDomain, err := repo.FindUserByEmail(EMAIL)

		assert.Nil(t, err)
		assert.NotNil(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.NotNil(t, userDomain.GetEmail(), userEntity.Email)
		assert.NotNil(t, userDomain.GetName(), userEntity.Name)
		assert.NotNil(t, userDomain.GetPassword(), userEntity.Password)

	})

	mtestDB.Run("when_sending_a_invalid_email_returns_error", func(mt *mtest.T) {

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		userDomain, err := repo.FindUserByEmail(EMAIL + "_erro")

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})

	mtestDB.Run("when_sending_an_email_absent", func(mt *mtest.T) {

		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, colleactionName), mtest.FirstBatch))

		databasemock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databasemock)
		userDomain, err := repo.FindUserByEmail(EMAIL)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	type fields struct {
		databaseConnection *mongo.Database
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   model.UserDomainInterface
		want1  *rest_err.RestErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ur := &userRepository{
				databaseConnection: tt.fields.databaseConnection,
			}
			got, got1 := ur.FindUserByEmailAndPassword(tt.args.email, tt.args.password)
			assert.Equalf(t, tt.want, got, "FindUserByEmailAndPassword(%v, %v)", tt.args.email, tt.args.password)
			assert.Equalf(t, tt.want1, got1, "FindUserByEmailAndPassword(%v, %v)", tt.args.email, tt.args.password)
		})
	}
}
