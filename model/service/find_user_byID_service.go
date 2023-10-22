package service

import (
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

func (ud *userDomainService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init FindUserByIDService ")
	return ud.userRepository.FindUserByID(id)
}
