package controller

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserControllerInterface_FindUserByEmail(t *testing.T) {

	controller := NewUserControllerInterface(nil)

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
