package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var GlobalMongoClient *mongo.Client
var mongoErr error

func InitializeMongoDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://432satyam432:vzImdy4NxdHpDIEB@cluster0.jqknrbi.mongodb.net/?appName=Cluster0").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	GlobalMongoClient, mongoErr = mongo.Connect(context.TODO(), opts)
	if mongoErr != nil {
		fmt.Println("Error Initializing Mongo Client")
	}

	// defer func() {
	// 	if mongoErr = GlobalMongoClient.Disconnect(context.TODO()); mongoErr != nil {
	// 		panic(mongoErr)
	// 	}
	// }()
}

func GetMongoCLient() *mongo.Client {
	if GlobalMongoClient == nil {
		InitializeMongoDB()
	}
	return GlobalMongoClient
}
