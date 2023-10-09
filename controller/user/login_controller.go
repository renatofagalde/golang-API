package controller

import (
	"github.com/gin-gonic/gin"
	"golang-basic/config/logger"
	"golang-basic/config/validation"
	"golang-basic/controller/model/request"
	"golang-basic/model"
	"golang-basic/view"
	"net/http"
)

func (uc *userControllerInterface) Login(c *gin.Context) {
	logger.Info("init login userController")

	var userRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		logger.Error("Erro ao validar user", errRest)
		c.JSON(errRest.Code, errRest)
		return
	}
	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
	domainResult, token, err := uc.service.LoginService(domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o login ", err)
		return
	}
	logger.Info("init login userController")
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
