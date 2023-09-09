package controller

import (
	"github.com/gin-gonic/gin"
	"golang-basic/config/logger"
	"golang-basic/config/validation"
	"golang-basic/controller/model/request"
	"net/http"
)

func Create(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		logger.Error("Erro ao validar user", errRest)
		c.JSON(errRest.Code, errRest)
		return
	}
	logger.Info("criando usuário %s")
	c.JSON(http.StatusOK, userRequest)
}
