package service

import (
	"golang-API/config/logger"
	"golang-API/config/rest_err"
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
