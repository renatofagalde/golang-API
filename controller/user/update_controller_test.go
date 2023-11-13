package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"golang-API/config/rest_err"
	"golang-API/controller/model/request"
	"golang-API/model"
	"golang-API/test/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserControllerInterface_DeleteUser(t *testing.T) {

	crtl := gomock.NewController(t)
	defer crtl.Finish()

	service := mocks.NewMockUserDomainService(crtl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_body_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserUpdateRequest{
			Name: "",
			Age:  -1,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, "PUT", url.Values{}, []gin.Param{}, stringReader)
		controller.Update(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_userId_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserUpdateRequest{
			Name: "teste user",
			Age:  10,
		}

		param := []gin.Param{
			{
				Key:   "id",
				Value: "test",
			},
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, "PUT", url.Values{}, param, stringReader)
		controller.Update(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("object_is_valid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "id",
				Value: id,
			},
		}

		userRequest := request.UserUpdateRequest{
			Name: "Test User test",
			Age:  10,
		}

		domain := model.NewUseUpdaterDomain(
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UpdateService(id, domain).Return(rest_err.NewInternalServerError("error test"))

		MakeRequest(context, "PUT", url.Values{}, param, stringReader)
		controller.Update(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("object_is_valid_and_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "id",
				Value: id,
			},
		}

		userRequest := request.UserUpdateRequest{
			Name: "Test User test",
			Age:  10,
		}

		domain := model.NewUseUpdaterDomain(
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UpdateService(id, domain).Return(nil)

		MakeRequest(context, "PUT", url.Values{}, param, stringReader)
		controller.Update(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

}
