package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-API/config/rest_err"
	"golang-API/model"
)

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPassword(email string, password string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(id string, domainInterface model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(id string) *rest_err.RestErr
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}
