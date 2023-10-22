package service

import (
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

func (ud *userDomainService) findUserByEmailAndPasswordService(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init findUserByEmailAndPasswordService")
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
