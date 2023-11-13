package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang-API/config/logger"
	"golang-API/config/rest_err"
	"golang-API/config/validation"
	"golang-API/controller/model/request"
	"golang-API/model"
	"net/http"
	"strings"
)

func (uc *userControllerInterface) Update(c *gin.Context) {
	logger.Info("init update userController")

	var userRequest request.UserUpdateRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(id) == "" {
		errRest := validation.ValidateUserError(err)
		logger.Error("Erro ao validar user", errRest)
		c.JSON(errRest.Code, errRest)
		return
	}

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUseUpdaterDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateService(id, domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o update ", err)
		return
	}
	logger.Info("init update userController successfuly",
		zap.String("id", id))
	c.Status(http.StatusOK)
}
