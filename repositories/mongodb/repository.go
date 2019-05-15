package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client            *mongo.Client
	dialogsCollection *mongo.Collection
}

func New(uri string) (*Repository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	collection := client.Database("neocortex").Collection("dialogs")

	return &Repository{
		client:            client,
		dialogsCollection: collection,
	}, nil
}
