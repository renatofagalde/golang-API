package service

import (
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

func (ud *userDomainService) Update(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init update model", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(id, userDomain)
	if err != nil {
		logger.Error("Error chamando repositório", err, zap.String("journey", "updateUser"))
		return err
	}
	logger.Info("init update model sucessfuly", zap.String("journey", "updateUser"))
	return nil
}
