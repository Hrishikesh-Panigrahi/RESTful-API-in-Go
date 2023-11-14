package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// type User struct{
// 	name string `json:"Name"`
// }

func main() {
	fmt.Println("Welcome to go_mysql")
	db, err := sql.Open("mysql", "root:Hrishikesh@38@tcp(127.0.0.1:3306)/user.users" )

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert , err := db.Query("Insert into user.users values(2, 'rishi')");

	if(err!=nil){
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Successfully inserted to mysql database")
}