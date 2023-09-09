package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/validation"
	"golang-basic/controller/model/request"
	"golang-basic/model"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
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
	domain := model.NewUserDomain(userRequest.Email,
		userRequest.Password, userRequest.Name, userRequest.Age)
	if err := domain.Create(); err != nil {
		c.JSON(err.Code, err)
		return
	}
	logger.Info("init created userController")
	c.JSON(http.StatusOK, userRequest)
}
