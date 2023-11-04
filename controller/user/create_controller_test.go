package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang-basic/config/rest_err"
	"golang-basic/controller/model/request"
	"golang-basic/model"
	"golang-basic/test/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestUserControllerInterface_Create(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	service := mocks.NewMockUserDomainService(crtl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_length_password_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test_error@test.com",
			Password: "123",
			Name:     "test test",
			Age:      19,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, "POST", url.Values{}, nil, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})
	t.Run("validation_name_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test_error@test.com",
			Password: "1234567",
			Name:     "a",
			Age:      19,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, "POST", url.Values{}, nil, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})
	t.Run("validation_emailgot_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test_error",
			Password: "1234567",
			Name:     "test of name",
			Age:      19,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, "POST", url.Values{}, nil, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})
	t.Run("validation_emailgot_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test_error",
			Password: "1234567@",
			Name:     "test of name",
			Age:      19,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, "POST", url.Values{}, nil, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("object_is_valid_service_returns_error", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "1234567@",
			Name:     "test of name",
			Age:      18,
		}

		domain := model.NewUserDomain(userRequest.Email,
			userRequest.Password, userRequest.Name, userRequest.Age)
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().CreateService(domain).Return(nil, rest_err.NewInternalServerError("error test"))

		MakeRequest(context, "POST", url.Values{}, []gin.Param{}, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("object_is_valid_service_returns_success", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "1234567@",
			Name:     "test of name",
			Age:      18,
		}

		domain := model.NewUserDomain(userRequest.Email,
			userRequest.Password, userRequest.Name, userRequest.Age)
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().CreateService(domain).Return(domain, nil)

		MakeRequest(context, "POST", url.Values{}, []gin.Param{}, stringReader)
		controller.Create(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)

	})

}
