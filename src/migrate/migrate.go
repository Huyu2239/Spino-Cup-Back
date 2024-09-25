package main

import (
	"api/db"
	"api/model"
	"fmt"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&model.Quiz2{})

	if err != nil {
		log.Fatalln(err.Error())
	}
}
