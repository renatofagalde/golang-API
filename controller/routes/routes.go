package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-basic/controller/user"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/:id", userController.FindUserById)
	r.PUT("/:id", userController.Update)
	r.GET("/email/:email", userController.FindUserByEmail)
	r.POST("/", userController.Create)
	r.DELETE("/:id", userController.Delete)
	r.POST("/login", userController.Login)
}
