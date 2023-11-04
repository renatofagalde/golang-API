package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"golang-API/config/logger"
	"golang-API/config/rest_err"
	"os"
)

func (ur *userRepository) DeleteUser(id string) *rest_err.RestErr {

	logger.Info("init delete user repository", zap.String("journey", "deleteUser"))

	collection := ur.databaseConnection.Collection(os.Getenv("MONGO_DB_COLLECTION"))
	idHex, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: idHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
