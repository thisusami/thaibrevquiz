package db

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDbProperty struct {
	Client  *mongo.Client
	DB      string
	Timeout time.Duration
}

func newMongClient(mongoServerURL string) (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(mongoServerURL))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoDbProperty(mongoServerURL, mongoDb string) (*MongoDbProperty, error) {
	mongoClient, err := newMongClient(mongoServerURL)
	if err != nil {
		log.Println("Error connecting to MongoDB:", err)
		panic(err)
	}
	log.Println("Connected to MongoDB")
	repo := &MongoDbProperty{
		Client:  mongoClient,
		DB:      mongoDb,
		Timeout: 10 * time.Second,
	}
	return repo, nil

}
