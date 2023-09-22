package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
	"golang-basic/model/repository/entity/convert"
	"os"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("init create user repository", zap.String("journey", "createUser"))

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_DB_COLLECTION))
	value := convert.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)
	return convert.ConvertEntityToDomain(*value), nil
}
