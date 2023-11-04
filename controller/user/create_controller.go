package controller

import (
	"github.com/gin-gonic/gin"
	"golang-API/config/logger"
	"golang-API/config/validation"
	"golang-API/controller/model/request"
	"golang-API/model"
	"golang-API/view"
	"net/http"
)

func (uc *userControllerInterface) Create(c *gin.Context) {
	logger.Info("init create userController")

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		logger.Error("Erro ao validar user", errRest)
		c.JSON(errRest.Code, errRest)
		return
	}
	domain := model.NewUserDomain(userRequest.Email,
		userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.service.CreateService(domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o create ", err)
		return
	}
	logger.Info("init create userController")

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
