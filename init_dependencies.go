package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-basic/config/logger"
	controller "golang-basic/controller/user"
	"golang-basic/model/repository"
	"golang-basic/model/service"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	logger.Info("inicializacao das dependencias")
	repo := repository.NewUserRepository(database)
	userDomainService := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(userDomainService)
}
