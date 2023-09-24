package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang-basic/config/logger"
	"golang-basic/config/rest_err"
	"golang-basic/model"
	"golang-basic/model/repository/entity/convert"
	"os"
)

func (ur *userRepository) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("init update user repository", zap.String("journey", "updateUser"))

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_DB_COLLECTION))
	value := convert.ConvertDomainToEntity(userDomain)

	idHex, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: idHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
