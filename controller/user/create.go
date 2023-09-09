package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/validation"
	"golang-basic/controller/model/request"
	"net/http"
)

func Create(c *gin.Context) {
	logger.Info("init create userController", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		logger.Error("Erro ao validar user", errRest, zap.String("journey", "createUser"))
		c.JSON(errRest.Code, errRest)
		return
	}
	logger.Info("init created userController")
	c.JSON(http.StatusOK, userRequest)
}
