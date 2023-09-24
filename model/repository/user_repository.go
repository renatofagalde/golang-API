package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang-basic/config/rest_err"
	"golang-basic/model"
)

type userRepository struct {
	databaseConnection *mongo.Database
}
type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(id string, domainInterface model.UserDomainInterface) *rest_err.RestErr
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

const (
	MONGO_DB_COLLECTION = "MONGO_DB_COLLECTION"
)
