package models

import (
	"context"
	"log" 

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options" 
)

// Connclient holds connection handle
// It is best practice to keep a client that is connected to MongoDB around so that the application can make use of connection pooling.
var Connclient *mongo.Client

// CloseConnections will close all the connections in the end
// if your application no longer requires a connection, the connection can be closed with client.Disconnect()
func CloseConnections() {
	err := Connclient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connections to MongoDB closed.")
	}
}

//GetClient gives a MongoDB Client
func GetClient(connString string) {
	var err error
	clientOptions := options.Client().ApplyURI(connString)
	Connclient, err = mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = Connclient.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = Connclient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to MongoDB!")
	}
}
