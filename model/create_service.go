package model

import (
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
)

func (ud *UserDomain) Create() *rest_err.RestErr {
	logger.Info("init createUser model", zap.String("journey", "createUser"))
	ud.EncryptPassword()
	logger.Info(ud.Password)
	return nil
}
