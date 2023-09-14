package service

import (
	"fmt"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

func (ud *userDomainService) Create(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("init createUser model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	fmt.Println(fmt.Sprintf("Password %s", userDomain.GetPassword()))
	return nil
}
