package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/leonfaneite/backend/src/api_twitter"
	"github.com/leonfaneite/backend/src/dbpostgrest"
	"github.com/leonfaneite/backend/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

var userCollection = api_twitter.GetMongoDbConnection().Database("twitterdb").Collection("twitter_search")

// CreateUser create a user in the postgres db
func Create_Words(w http.ResponseWriter, r *http.Request) {

	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var word models.Query

	// decode the json request to words
	err := json.NewDecoder(r.Body).Decode(&word)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := Insert_words(word)

	res := response{
		ID:      insertID,
		Message: "Find created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func Insert_words(words models.Query) int {

	db := dbpostgrest.GetConnect()
	defer db.Close()
	sqlStatement := `INSERT INTO query(Id,Words) VALUES ($1,$2) RETURNING Id`

	// the inserted id will store in this id
	var Id int

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, words.Id, words.Words).Scan(&Id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", Id)

	return Id

	// return the inserted id

}

//
func Delet_Words(w http.ResponseWriter, r *http.Request) {

	collection := userCollection
	// Set header
	w.Header().Set("Content-Type", "application/json")

	filte := bson.D{{}}

	deleteResult, err := collection.DeleteMany(context.TODO(), filte)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Starship deleted (s)\n", deleteResult.DeletedCount)

	json.NewEncoder(w).Encode(deleteResult)

}

/////////////////////////////////////////////////////////////////7

func Get_all_Words(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var results []primitive.M                                   //slice for multiple documents
	cur, err := userCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}
