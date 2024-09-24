package main

import (
	"api/db"
	"api/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Quiz{}, &model.Score{}, &model.User{})
}
