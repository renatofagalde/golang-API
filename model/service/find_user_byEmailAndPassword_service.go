package service

import (
	"golang-API/config/logger"
	"golang-API/config/rest_err"
	"golang-API/model"
)

func (ud *userDomainService) findUserByEmailAndPasswordService(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init findUserByEmailAndPasswordService")
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
