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
	Create(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	Update(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
