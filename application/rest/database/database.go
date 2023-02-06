package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewConectionDatabase() (*mongo.Database, error) {
	uri := "mongodb+srv://admin:admin123@brahyan0504.civzx6e.mongodb.net/?retryWrites=true&w=majority"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}

	return client.Database("test"), nil
}
