package main

import (
	//"fmt
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"



)

func main() {
	//initialize connection pool
	db, err := gorm.Open(mysql.Open("root:Nn@2681999@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	//creating an instance of env containing the connection pool
	env := &Env{db:db}
	r := mux.NewRouter()
	//handler functions which are accessing the connection pool
	r.HandleFunc("/path", env.PostFilePath).Methods("POST")
	r.HandleFunc("/path/users", env.GetUserInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", r))

}
