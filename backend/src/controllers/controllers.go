package controllers

import (
	"fmt"
	"github.com/leonfaneite/backend/src/dbpostgrest"
	"net/http"
	"encoding/json"
	"github.com/leonfaneite/backend/src/models"
	"log"
	"database/sql"
	"github.com/gorilla/mux"
	"strconv"
	_ "github.com/lib/pq"
	
)

type response struct {
    ID      int  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

// CreateUser create a user in the postgres db
func CreateWord(w http.ResponseWriter, r *http.Request) {
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

    // format a response object
    res := response{
        ID:      insertID,
        Message: "Find created successfully",
    }

    // send the response
    json.NewEncoder(w).Encode(res)
}



func GetWord(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get the userid from the request params, key is "id"
    params := mux.Vars(r)

    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])

    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    // call the getUser function with user id to retrieve a single user
    find, err := Get_words(int(id))

    if err != nil {
        log.Fatalf("Unable to get user. %v", err)
    }

    // send the response
    json.NewEncoder(w).Encode(find)
}



/////////Add Order////////////
func Insert_words(words models.Query) int  {
	
	db := dbpostgrest.GetConnect()
	defer db.Close()
	sqlStatement := `INSERT INTO query(Id,Words) VALUES ($1,$2) RETURNING Id`

    // the inserted id will store in this id
    var Id int

    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement,words.Id, words.Words).Scan(&Id)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    fmt.Printf("Inserted a single record %v", Id)

    // return the inserted id
    return Id

}


func Get_words(id int) (models.Query, error) {
    // create the postgres db connection
	db := dbpostgrest.GetConnect()
	defer db.Close()

    // create a user of models.User type
    var word models.Query

    // create the select sql query
    sqlStatement := `SELECT * FROM query WHERE Id=$1`

    // execute the sql statement
    row := db.QueryRow(sqlStatement, id)

    // unmarshal the row object to user
    err := row.Scan(&word.Id, &word.Words)

    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return word, nil
    case nil:
        return word, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }

    // return empty user on error
    return word, err
}