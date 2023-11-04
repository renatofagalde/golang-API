package service

import (
	"golang-API/config/logger"
	"golang-API/config/rest_err"
	"golang-API/model"
)

func (ud *userDomainService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init FindUserByIDService ")
	return ud.userRepository.FindUserByID(id)
}
