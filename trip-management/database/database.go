package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client
var MongoDb *mongo.Database
var MongoCollection *MongoDatabaseCollection

type MongoDatabaseCollection struct {
	TripsCollection *mongo.Collection
}

func ConnectDatabase() *mongo.Client {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:27017/uber-trip-db"))
	if err != nil {
		panic(err)
	}
	//ping db
	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	MongoClient = mongoClient
	return mongoClient
}

func AccessDatabase(mongoClient *mongo.Client, dbName string) *mongo.Database {
	mongoDb := mongoClient.Database(dbName)
	MongoDb = mongoDb
	return mongoDb
}

func AccessTripsCollection(mongoDatabase *mongo.Database) {
	tripsCollection := mongoDatabase.Collection("trips")

	MongoCollection = &MongoDatabaseCollection{
		TripsCollection: tripsCollection,
	}
}
