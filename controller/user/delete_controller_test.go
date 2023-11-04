package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"golang-basic/config/rest_err"
	"golang-basic/model"
	"golang-basic/test/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDeleteController_DeleteUser(t *testing.T) {

	crtl := gomock.NewController(t)
	defer crtl.Finish()

	service := mocks.NewMockUserDomainService(crtl)
	controller := NewUserControllerInterface(service)

	t.Run("id_is_invalid_returns_error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		params := []gin.Param{
			{
				Key:   "id",
				Value: "test_at_erro",
			},
		}

		MakeRequest(context, "DELETE", url.Values{}, params, nil)
		controller.Delete(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("id_is_valid_service_return_error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		ID := primitive.NewObjectID().Hex()
		params := []gin.Param{
			{
				Key:   "id",
				Value: ID,
			},
		}

		service.EXPECT().FindUserByIDService(ID).Return(nil,
			rest_err.NewInternalServerError("error test"))

		MakeRequest(context, "GET", url.Values{}, params, nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("return_success", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		ID := primitive.NewObjectID().Hex()
		params := []gin.Param{
			{
				Key:   "id",
				Value: ID,
			},
		}

		service.EXPECT().FindUserByIDService(ID).
			Return(model.NewUserDomain("test@test.com", "pass", "name", 18), nil)

		MakeRequest(context, "GET", url.Values{}, params, nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})
}
