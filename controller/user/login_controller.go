package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/validation"
	"golang-basic/controller/model/request"
	"golang-basic/model"
	"golang-basic/view"
	"net/http"
)

func (uc *userControllerInterface) Login(c *gin.Context) {
	logger.Info("init login userController", zap.String("journey", "loginUser"))

	var userRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)
		logger.Error("Erro ao validar user", errRest, zap.String("journey", "loginUser"))
		c.JSON(errRest.Code, errRest)
		return
	}
	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
	domainResult, err := uc.service.LoginService(domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o login ", err, zap.String("journey", "loginUser"))
		return
	}
	logger.Info("init login userController", zap.String("journey", "loginUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
