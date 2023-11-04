package service

import (
	"golang-API/config/logger"
	"golang-API/config/rest_err"
	"golang-API/model"
)

func (ud *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init FindUserByEmailService ")
	return ud.userRepository.FindUserByEmail(email)
}
