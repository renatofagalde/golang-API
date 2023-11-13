package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	controller "golang-API/controller/user"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
)

//func TestMain(m *testing.M) {
//
//	os.Setenv("MONGO_DB_COLLECTION", "db_collection")
//	os.Setenv("MYSQL_DB_USER", "test_user")
//
//	closeConnection := func() {}
//	Database, closeConnection = mongodb.OpenConnection()
//
//	repo := repository.NewUserRepository(Database)
//	userService := service2.NewUserDomainService(repo)
//	UserController = controller.NewUserControllerInterface(userService)
//
//	defer func() {
//		os.Clearenv()
//		closeConnection()
//	}()
//
//	os.Exit(m.Run())
//}

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
