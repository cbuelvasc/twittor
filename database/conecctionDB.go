package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConecctiont = Connect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://twittor_user_admin:admin1234@cluster0.boumn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

func Connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Printf("Connected successfully...!")
	return client
}

func CheckConnection() int {
	err := MongoConecctiont.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	log.Printf("Connected successfully...!")
	return 1
}
