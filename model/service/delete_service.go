package service

import (
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
)

func (ud *userDomainService) DeleteService(id string) *rest_err.RestErr {
	logger.Info("init delete model")

	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("Error chamando repositório", err)
		return err
	}
	logger.Info("init delete model sucessfuly")
	return nil
}
