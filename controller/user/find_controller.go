package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-API/config/logger"
	"golang-API/config/rest_err"
	"golang-API/view"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {

	logger.Info("init FindUserByID find_controller")
	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		message := "ID não é válido"
		logger.Error(message, err)
		errorMessage := rest_err.NewBadRequestError(message)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDService(id)
	if err != nil {
		c.JSON(err.Code, err)
		message := "Erro ao recuperar ID"
		logger.Error(message, err)
		return
	}
	logger.Info("init FindUserByID find_controller successfuly")

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("init FindUserByEmail find_controller")
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
		logger.Error(message, err)
		c.JSON(err.Code, err)
		return
	}
	logger.Info("init FindUserByEmail find_controller successfuly")
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
