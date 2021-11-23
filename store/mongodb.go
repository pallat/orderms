package store

import (
	"context"

	"github.com/pallat/micro/order"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStore struct {
	*mongo.Collection
}

func NewMongoDBStore(dsn string) *MongoDBStore {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database("myapp").Collection("orders")

	return &MongoDBStore{Collection: collection}
}

func (s *MongoDBStore) Save(order order.Order) error {
	_, err := s.Collection.InsertOne(context.Background(), order)
	return err
}
