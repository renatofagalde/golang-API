package service

import (
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

func (ud *userDomainService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init FindUserByIDService ", zap.String("journey", "findUserByID"))
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init FindUserByEmailService ", zap.String("journey", "FindUserByEmailService"))
	return ud.userRepository.FindUserByEmail(email)
}
func (ud *userDomainService) findUserByEmailAndPasswordService(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init findUserByEmailAndPasswordService ", zap.String("journey", "login"))
	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
