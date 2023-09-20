package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	MONGO_DB_URL      = "MONGO_DB_URL"
	MONGO_DB_DATABASE = "MONGO_DB_DATABASE"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	mongodbUri := os.Getenv(MONGO_DB_URL)
	mongodbDatabase := os.Getenv(MONGO_DB_DATABASE)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodbDatabase), nil
}
