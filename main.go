package main

import (
	//"fmt
	"log"
	"github.com/gorilla/mux"
	"net/http"
	
)
func main() {

	InitailizeDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/path", PostFilePath).Methods("POST")
	r.HandleFunc("/path/users", GetUserInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", r))

}
