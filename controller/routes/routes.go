package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang-basic/controller/user"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/:id", controller.FindUserById)
	r.GET("/email/:email", controller.FindUserByEmail)
	r.POST("/", controller.Create)
	r.PUT("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
}
