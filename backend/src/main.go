package main

import (
	

	"github.com/leonfaneite/backend/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
	
	
	"fmt"


)

func main(){

	var port string = "3000"
	

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/find/{id}", controllers.GetWord).Methods("GET")
	apiRouter.HandleFunc("/add", controllers.CreateWord).Methods("POST")
	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)

}
