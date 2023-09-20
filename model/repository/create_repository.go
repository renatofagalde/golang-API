package repository

import (
	"context"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
	"os"
)

const (
	MONGO_DB_COLLECTION = "MONGO_DB_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init create user repository")

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_DB_COLLECTION))
	value, err := userDomain.ToJSON()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	userDomain.AtribuirID(result.InsertedID.(string)) //cast para string
	return userDomain, nil
}
