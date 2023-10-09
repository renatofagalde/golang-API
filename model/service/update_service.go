package service

import (
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

func (ud *userDomainService) UpdateService(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init update model")

	err := ud.userRepository.UpdateUser(id, userDomain)
	if err != nil {
		logger.Error("Error chamando repositório", err)
		return err
	}
	logger.Info("init update model sucessfuly")
	return nil
}
