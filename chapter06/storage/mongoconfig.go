package storage

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoStorage implementa nossa interface de armazenamento
type MongoStorage struct {
	*mongo.Client
	DB string
	Collection string
}

//NewMongoStorage inicializa um MongoStorage
func NewMongoStorage(ctx context.Context, connection, db, collection string) (*MongoStorage, error)  {
	
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connection))
	if err != nil {
		return nil, err
	}

	ms := MongoStorage {
		Client: client,
		DB: db,
		Collection: collection,
	}

	return &ms, nil
}