package tests

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	controller "golang-API/controller/user"
	"golang-API/model/repository"
	service2 "golang-API/model/service"
	"golang-API/test/mongodb"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Setenv("MONGO_DB_COLLECTION", "db_collection")
	os.Setenv("MYSQL_DB_USER", "test_user")

	closeConnection := func() {}
	Database, closeConnection = mongodb.OpenConnection()

	repo := repository.NewUserRepository(Database)
	userService := service2.NewUserDomainService(repo)
	UserController = controller.NewUserControllerInterface(userService)

	defer func() {
		os.Clearenv()
		closeConnection()
	}()

	os.Exit(m.Run())
}
func TestFindUserByEmail(t *testing.T) {

	t.Run("user_not_found_with_this_email", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		params := []gin.Param{
			{
				Key:   "email",
				Value: "test@test.com",
			},
		}

		MakeRequest(context, "GET", url.Values{}, params, nil)
		UserController.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	})

	t.Run("user_found_with_this_email", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		ID := primitive.NewObjectID().Hex()

		_, err := Database.
			Collection(os.Getenv("MONGO_DB_COLLECTION")).
			InsertOne(context.Background(), bson.M{"_id": ID, "name": t.Name(), "email": "test@test.com"})

		if err != nil {
			t.Fatal(err)
			return
		}

		params := []gin.Param{
			{
				Key:   "email",
				Value: "test@test.com",
			},
		}

		MakeRequest(ctx, "GET", url.Values{}, params, nil)
		UserController.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})
}

func TestFindUserByID(t *testing.T) {

	t.Run("user_not_found_with_this_ID", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		ID := primitive.NewObjectID().Hex()

		params := []gin.Param{
			{
				Key:   "id",
				Value: ID,
			},
		}

		MakeRequest(context, "GET", url.Values{}, params, nil)
		UserController.FindUserByID(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	})

	t.Run("user_found_with_this_ID", func(t *testing.T) {

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

		MakeRequest(ctx, "GET", url.Values{}, params, nil)
		UserController.FindUserByID(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})
}
