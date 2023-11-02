package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func (uc *userControllerInterface) FindUserById(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := GetTestGinContext(recorder)
	params := []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	u := url.Values{}
	u.Set("foo", "bar")

	MakeFunction(context, params, u)
	GetUserId(context)

	//t.Errorf("Received %v (type %v), expected %v (type %v)", recorder.Code, reflect.TypeOf(recorder.Code),
	//	http.StatusOK, reflect.TypeOf(http.StatusOK))
	assert.EqualValues(t, http.StatusOK, recorder.Code)
	got, _ := strconv.Atoi(recorder.Body.String())
	assert.EqualValues(t, 1, got)

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

func MakeFunction(c *gin.Context, param gin.Params, u url.Values, body io.ReadCloser) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
