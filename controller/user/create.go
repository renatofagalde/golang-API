package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	"golang-basic/config/logger"
	"golang-basic/config/validation"
	"golang-basic/controller/model/request"
	"net/http"
)

func Create(c *gin.Context) {
	var userRequest request.UserRequest

	logger.Info("init create userController", zapcore.Field{
		Key:    "journey",
		String: "createUser",
	})

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		logger.Error("Erro ao validar user", errRest, zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})
		c.JSON(errRest.Code, errRest)
		return
	}
	logger.Info("init created userController", zapcore.Field{
		Key:    "journey",
		String: "createUser",
	})
	c.JSON(http.StatusOK, userRequest)
}
