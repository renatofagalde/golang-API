package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

func TestUserControllerInterface_Login(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	service := mocks.NewMockUserDomainService(crtl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_length_password_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLoginRequest{
			Email:    "test_error@test.com",
			Password: "123",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, "POST", url.Values{}, nil, stringReader)
		controller.Login(context)

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
		}

		domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginService(domain).Return(nil, "", rest_err.NewInternalServerError("error test"))

		MakeRequest(context, "POST", url.Values{}, []gin.Param{}, stringReader)
		controller.Login(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

	})

	t.Run("object_is_valid_service_returns_success", func(t *testing.T) {

		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		token := uuid.NewString()

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "1234567@",
		}
		domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().LoginService(domain).Return(domain, token, nil)

		MakeRequest(context, "POST", url.Values{}, []gin.Param{}, stringReader)
		controller.Login(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, recorder.Header().Values("Authorization")[0], token)

	})

}
