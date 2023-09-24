package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"net/http"
)

func (uc *userControllerInterface) Delete(c *gin.Context) {
	logger.Info("init delete userController", zap.String("journey", "deleteUser"))

	id := c.Param("id")

	err := uc.service.DeleteService(id)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o delete ", err, zap.String("journey", "deleteUser"))
		return
	}
	logger.Info("init delete userController successfuly",
		zap.String("id", id),
		zap.String("journey", "deleteUser"))
	c.Status(http.StatusOK)
}
