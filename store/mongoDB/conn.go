package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection(databaseURL string) (*mongo.Database, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		return nil, err
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("priceService")

	return db, nil
}
