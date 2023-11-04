package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-API/config/logger"
	controller "golang-API/controller/user"
	"golang-API/model/repository"
	"golang-API/model/service"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	logger.Info("inicializacao das dependencias")
	repo := repository.NewUserRepository(database)
	userDomainService := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(userDomainService)
}
