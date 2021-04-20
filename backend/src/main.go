package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leonfaneite/backend/src/controllers"
)

func main() {

	var port string = "3000"

	fmt.Printf("Server running at port 3000")

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/find", controllers.Get_all_Words).Methods("GET")
	apiRouter.HandleFunc("/add", controllers.Create_Words).Methods("POST")
	apiRouter.HandleFunc("/delete", controllers.Delet_Words_mongo).Methods("DELETE")

	err := http.ListenAndServe(":"+port, router)
	//fmt.Printf("Server running at port %s", port)

	if err != nil {
		panic(err)
	}

}
