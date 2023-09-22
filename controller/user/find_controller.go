package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/view"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {

	logger.Info("init FindUserById find_controller", zap.String("journey", "FindUserById"))
	id := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		message := "ID não é válido"
		logger.Error(message, err, zap.String("journey", "FindUserById"))
		errorMessage := rest_err.NewBadRequestError(message)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByIDService(id)
	if err != nil {
		c.JSON(err.Code, err)
		message := "Erro ao recuperar ID"
		logger.Error(message, err, zap.String("journey", "FindUserById"))
		return
	}
	logger.Info("init FindUserById find_controller successfuly", zap.String("journey", "FindUserById"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("init FindUserByEmail find_controller", zap.String("journey", "FindUserByEmail"))
	email := c.Param("email")
	if _, err := mail.ParseAddress(email); err != nil {
		message := "Email não é válido"
		errorMessage := rest_err.NewBadRequestError(message)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByEmailService(email)
	if err != nil {
		message := "Erro ao recuperar email"
		logger.Error(message, err, zap.String("journey", "FindUserByEmail"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("init FindUserByEmail find_controller successfuly", zap.String("journey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
