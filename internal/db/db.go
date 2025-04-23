package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var mongoOnce sync.Once

var clientInstanceerror error

type Collection string

const (
	ProductsCollection Collection = "products"
	// Can make the same with user e.g.
)

const (
	url = "mongodb://root:example@localhost:27017/"
	// // If inside of docker compose as a service use mongodb://root:example@mongo:27017/
	Database = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(url)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		clientInstance = client
		clientInstanceerror = err
	})

	return clientInstance, clientInstanceerror

}

// * reference the value

// & point the block of memory that stores this object (Hex address)

//var x int = 10
