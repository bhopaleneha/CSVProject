package main

import (
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
)

func ReadCsvFile(csv_file *os.File) [][]string {
	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return records
}

var MapUserId = make(map[string]struct{})

func ListValidUsers(records [][]string) []User {

	var ListOfValidUser []User

	for _, line := range records {
		Active, _ := strconv.ParseBool(line[4])
		user := User{
			Id:       line[0],
			Name:     line[1],
			Email:    line[2],
			Phone:    line[3],
			IsActive: Active,
		}

		isValidUser := Validate(&user)
		if isValidUser {
			ListOfValidUser = append(ListOfValidUser, user)
			MapUserId[user.Id] = struct{}{}
		}

	}
	return ListOfValidUser
}

var standardLogger = NewLogger()			// to handle error

func Validate(user *User) bool {
	Validations := true
	if user.Id == "" {
		id := uuid.New()
		user.Id = id.String()
	}

	if user.Name == "" {
		standardLogger.InvalidArgValue(" name as blank field", user.Id)
		Validations = false
	}
	if user.Email == "" {
		standardLogger.InvalidArgValue("email as a blank field", user.Id)
		Validations = false
	}
	if len(user.Phone) != 10 {
		standardLogger.InvalidArgValue("Phone Number as a ", user.Id)
		Validations = false
	}
	_, isUserAlreadyExist := MapUserId[user.Id]
	if isUserAlreadyExist {
		standardLogger.InvalidArgValue("existing user ", user.Id)
		Validations = false

	}
	return Validations
}
