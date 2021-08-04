package main

import (
	"encoding/json"
	"net/http"
	"os"
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
	
	
)

//Env Structure holds a connection pool
type Env struct{
	db *gorm.DB
}

//here we access our connection pool directly in handler
func (env *Env)createUsersInDatabase(users []User){
	for _,user :=range(users){
		if env.db.Model(&user).Where("id=?",user.Id).Updates(&user).RowsAffected==0{
			env.db.Create(&user)
		}
	}
	
}
func(env *Env) PostFilePath(w http.ResponseWriter, r *http.Request){
	var file File
	json.NewDecoder(r.Body).Decode(&file)
	csv_file, err := os.Open(file.PathOfFile)
	if err != nil {
			logrus.Fatal("unable to open csv file",err)
	}else{
		logrus.Info("File with posted path is available")
	}
	defer csv_file.Close()
	rec:=ReadCsvFile(csv_file)
	users:=ListValidUsers(rec)
	env.createUsersInDatabase(users)
	conToJson(users)

}


func(env *Env) GetUserInfo(w http.ResponseWriter, r *http.Request){

	var user []User
	env.db.Find(&user)
	json.NewEncoder(w).Encode(user)
}

