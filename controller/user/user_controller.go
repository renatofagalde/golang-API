package controller

import (
	"github.com/gin-gonic/gin"
	"golang-API/model/service"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{service: serviceInterface}
}

type UserControllerInterface interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	Update(c *gin.Context)

	Login(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
