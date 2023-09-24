package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/validation"
	"golang-basic/controller/model/request"
	"golang-basic/model"
	"net/http"
	"strings"
)

func (uc *userControllerInterface) Update(c *gin.Context) {
	logger.Info("init update userController", zap.String("journey", "updateUser"))

	var userRequest request.UserUpdateRequest
	id := c.Param("id")

	if err := c.ShouldBindJSON(&userRequest); err != nil || strings.TrimSpace(id) == "" {
		errRest := validation.ValidateUserError(err)
		logger.Error("Erro ao validar user", errRest, zap.String("journey", "updateUser"))
		c.JSON(errRest.Code, errRest)
		return
	}
	domain := model.NewUseUpdaterDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateService(id, domain)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o update ", err, zap.String("journey", "updateUser"))
		return
	}
	logger.Info("init update userController successfuly",
		zap.String("id", id),
		zap.String("journey", "updateUser"))
	c.Status(http.StatusOK)
}
