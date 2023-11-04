package controller

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"golang-API/config/rest_err"
	"golang-API/model"
	"golang-API/test/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserControllerInterface_FindUserByEmail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)

	controller := NewUserControllerInterface(service)

	t.Run("email_is_invalid_returns_error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		params := []gin.Param{
			{
				Key:   "email",
				Value: "test_at_erro",
			},
		}

		MakeRequest(context, "GET", url.Values{}, params, nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

	})

	t.Run("email_is_valid_service_return_error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		params := []gin.Param{
			{
				Key:   "email",
				Value: "test@test.com",
			},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").Return(nil,
			rest_err.NewInternalServerError("error test"))

		MakeRequest(context, "GET", url.Values{}, params, nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("return_success", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		params := []gin.Param{
			{
				Key:   "email",
				Value: "test@test.com",
			},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").
			Return(model.NewUserDomain("test@test.com", "pass", "name", 18), nil)

		MakeRequest(context, "GET", url.Values{}, params, nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})

}

func TestUserControllerInterface_FindUserByID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)

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

		MakeRequest(context, "GET", url.Values{}, params, nil)
		controller.FindUserByID(context)

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

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

func MakeRequest(c *gin.Context, method string, u url.Values, param gin.Params, body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("X-Request-ID", uuid.NewString())
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
