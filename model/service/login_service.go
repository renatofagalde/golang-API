package service

import (
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

func (ud *userDomainService) LoginService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("init loginUser model")

	userDomain.EncryptPassword()
	user, err := ud.findUserByEmailAndPasswordService(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		logger.Error("init loginUser erro ao validar", err)
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
