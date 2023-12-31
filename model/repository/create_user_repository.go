package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang-API/config/logger"
	"golang-API/config/rest_err"
	"golang-API/model"
	"golang-API/model/repository/entity/convert"
	"os"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init create user repository", zap.String("journey", "createUser"))

	collection := ur.databaseConnection.Collection(os.Getenv("MONGO_DB_COLLECTION"))
	value := convert.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)
	logger.Info(fmt.Sprintf("init create user -> %#v <- repository success", value), zap.String("journey", "createUser"))
	return convert.ConvertEntityToDomain(*value), nil
}
