package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
)

//structure for all fields in the database
type User struct {
	gorm.Model
	Id       string `json:"id" gorm:"primary_key"`
	Name     string `json :"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone_number"`
	IsActive bool   `json:"active_status"`
}

//structure to get a path of the file
type File struct {
	PathOfFile string `json:"PathOfFile"`
}

func conToJson(ListOfValidUser []User) {

	jsondata, err := json.MarshalIndent(ListOfValidUser, "", "  ") // convert to JSON

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Information of valid users in the json format ")
	fmt.Println(string(jsondata))
	jsonFile, err := os.Create("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	jsonFile.Write(jsondata)

}
