package api_twitter

import (

	//"database/sql"
	"fmt"

	"context"
	"log"

	_ "github.com/lib/pq"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	user = "root"

	pwd = "1234"

	port = 27017

	host = "localhost"

	//database = "twitterdb"
	//fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pwd, host, port)

	// mongo
)

func GetMongoDbConnection() *mongo.Client {

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, pwd, host, port))
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")
	return client

}
