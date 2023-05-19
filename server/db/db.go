package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB() *mongo.Collection {


	// Set client options

	clientOptions := options.Client().ApplyURI("mongodb+srv://amantarar01:"+os.Getenv("MongoPass")+"@cluster0.kf61u4b.mongodb.net/?retryWrites=true&w=majority")

	fmt.Println("clientOptions",clientOptions)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	fmt.Println("client",client)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// collection := client.Database("go_rest_api").Collection("books")

	collection := client.Database(os.Getenv("Database")).Collection(os.Getenv("Collection"))

	return collection
}

