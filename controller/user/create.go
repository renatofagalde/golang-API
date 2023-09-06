package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-basic/config/rest_err"
	"golang-basic/controller/model/request"
	"net/http"
)

func Create(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
	c.JSON(http.StatusOK, userRequest)
}
