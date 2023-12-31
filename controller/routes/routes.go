package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-API/controller/user"
	"golang-API/model"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/:id", model.VerifyTokenMiddleware, userController.FindUserByID)
	r.PUT("/:id", model.VerifyTokenMiddleware, userController.Update)
	r.GET("/email/:email", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/", userController.Create)
	r.DELETE("/:id", model.VerifyTokenMiddleware, userController.Delete)
	r.POST("/login", userController.Login)
}
