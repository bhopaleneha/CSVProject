# CSVProject

## Go Files With Respective Responsibilities
 1. ApiHandler: 
            PostFilePath function will open the csv file from the path we have specified using post method .
            Using gorm we will initialize the database . Transfer the entries of valid users  to the database 
            GetUserInfo will help to retrive the information from the database
2.  Validator:
            It will read the csv files and validate the entries present in the file
3.  Converter : 
         It has user struct having required fields for user information.Responsibility of converter is to convert data of valid users into json format 
4. main:
      It has main function which will iniatize the route using mux and apply get and post requests on respective function
 
               
