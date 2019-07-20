package mongodb

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client      *mongo.Client
	dialogs     *mongo.Collection
	views       *mongo.Collection
	actions     *mongo.Collection
	collections *mongo.Collection
}

type collection struct {
	Box    string   `json:"box"`
	Values []string `json:"values"`
}

type action struct {
	Name string            `json:"name"`
	Vars map[string]string `json:"vars"`
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

	dialogs := client.Database("neocortex").Collection("dialogs")
	views := client.Database("neocortex").Collection("views")
	actions := client.Database("neocortex").Collection("actions")
	collections := client.Database("neocortex").Collection("collections")

	// * Creating different 'boxes' for intents, entities, dialog nodes and context variables

	coll := new(collection)
	if err := collections.FindOne(context.Background(), bson.M{"box": "intents"}).Decode(coll); err != nil {
		_, err := collections.InsertOne(context.Background(), collection{Box: "intents", Values: []string{}})
		if err != nil {
			return nil, err
		}
	}

	if err := collections.FindOne(context.Background(), bson.M{"box": "entities"}).Decode(coll); err != nil {
		_, err := collections.InsertOne(context.Background(), collection{Box: "entities", Values: []string{}})
		if err != nil {
			return nil, err
		}
	}

	if err := collections.FindOne(context.Background(), bson.M{"box": "nodes"}).Decode(coll); err != nil {
		_, err := collections.InsertOne(context.Background(), collection{Box: "nodes", Values: []string{}})
		if err != nil {
			return nil, err
		}
	}

	if err := collections.FindOne(context.Background(), bson.M{"box": "context_vars"}).Decode(coll); err != nil {
		_, err := collections.InsertOne(context.Background(), collection{Box: "context_vars", Values: []string{}})
		if err != nil {
			return nil, err
		}
	}

	act := new(action)
	if err := actions.FindOne(context.Background(), bson.M{"name": "envs"}).Decode(act); err != nil {
		_, err := actions.InsertOne(context.Background(), action{Name: "envs", Vars: map[string]string{}})
		if err != nil {
			return nil, err
		}
	}

	return &Repository{
		client:      client,
		dialogs:     dialogs,
		views:       views,
		actions:     actions,
		collections: collections,
	}, nil
}
