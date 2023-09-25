package service

import (
	"golang-basic/config/rest_err"
	"golang-basic/model"
	"golang-basic/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}
type UserDomainService interface {
	CreateService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateService(string, model.UserDomainInterface) *rest_err.RestErr
	DeleteService(string) *rest_err.RestErr

	LoginService(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}
