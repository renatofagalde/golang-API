package service

import (
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}
type UserDomainService interface {
	Create(model.UserDomainInterface) *rest_err.RestErr
	Update(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
