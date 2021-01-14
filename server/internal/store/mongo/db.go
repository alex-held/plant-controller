package mongo

import (
	"context"
	"log"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	"plant-controller/internal/config"
)

type MongoDB struct {
	Database *mongo.Database
	Client   *mongo.Client
}

func Dial() (db *MongoDB, err error) {
	cfg := config.Get()
	
	if cfg.MongoURL == "" {
		return nil, nil
	}
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURL))
	if err != nil {
		log.Fatal(err)
	}
	
	return &MongoDB{Client: client, Database: client.Database(cfg.MongoDatabase)}, nil
}
