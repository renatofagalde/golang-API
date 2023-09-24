package service

import (
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
)

func (ud *userDomainService) DeleteService(id string) *rest_err.RestErr {
	logger.Info("init delete model", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(id)
	if err != nil {
		logger.Error("Error chamando repositório", err, zap.String("journey", "deleteUser"))
		return err
	}
	logger.Info("init delete model sucessfuly", zap.String("journey", "deleteUser"))
	return nil
}
