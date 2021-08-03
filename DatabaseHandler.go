package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

var db *gorm.DB
var err error
const DNS = "root:Nn@2681999@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"


func InitailizeDatabase(){
	db, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("cannot connect to db")
	}
	/*db.Exec("CREATE DATABASE users")
    db.Exec("USE users")
    db.AutoMigrate(&User{})*/
}
func createUsersInDatabase(users []User){
	for _,user :=range(users){
		if db.Model(&user).Where("id=?",user.Id).Updates(&user).RowsAffected==0{
			db.Create(&user)
		}
	}
	
}
func PostFilePath(w http.ResponseWriter, r *http.Request){
	var file File
	json.NewDecoder(r.Body).Decode(&file)
	csv_file, err := os.Open(file.PathOfFile)
	if err != nil {
			log.Fatal("unable to open csv file",err)
	}else{
		log.Info("File with posted path is available")
	}
	defer csv_file.Close()
	rec:=ReadCsvFile(csv_file)
	users:=ListValidUsers(rec)
	createUsersInDatabase(users)
	conToJson(users)

}


func GetUserInfo(w http.ResponseWriter, r *http.Request){

	var user []User
	db.Find(&user)
	json.NewEncoder(w).Encode(user)
}

