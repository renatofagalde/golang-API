package service

import (
	"golang-API/config/logger"
	"golang-API/config/rest_err"
	"golang-API/model"
)

func (ud *userDomainService) CreateService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init createUser model")

	user, _ := ud.FindUserByEmailService(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("email já existente")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	return userDomainRepository, nil
}
