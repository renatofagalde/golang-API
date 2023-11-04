package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"net/http"
)

func (uc *userControllerInterface) Delete(c *gin.Context) {
	logger.Info("init delete userController")

	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid ID")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteService(id)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Erro ao chamar o delete ", err)
		return
	}
	logger.Info("init delete userController successfuly",
		zap.String("id", id))
	c.Status(http.StatusOK)
}
