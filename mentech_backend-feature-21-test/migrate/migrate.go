package main

import (
	"fmt"
	"mentech-backend/db"
	"mentech-backend/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
