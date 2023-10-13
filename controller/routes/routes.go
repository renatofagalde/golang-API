package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-basic/controller/user"
	"golang-basic/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/:id", model.VerifyTokenMiddleware, userController.FindUserById)
	r.PUT("/:id", model.VerifyTokenMiddleware, userController.Update)
	r.GET("/email/:email", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/", userController.Create)
	r.DELETE("/:id", model.VerifyTokenMiddleware, userController.Delete)
	r.POST("/login", userController.Login)
}
