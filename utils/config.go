package utils

import (
	"fmt"
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/users"
	"log"
)

func CheckForEmptyDatabase() {
	db := database.GetSqlReadDB()
	rows, err := db.Query("show tables;")
	if err != nil {
		log.Fatal("Unable to retrieve tables ;(")
	}

	areThereTables := rows.Next()
	if !areThereTables {
		fmt.Println("No tables found. ")
		fmt.Print("Creating..")
		setupDatabaseTables()
		fmt.Println("..Finished.")
	}
}

func setupDatabaseTables() {
	db := database.GetSqlWriteDB()
	db.Exec(users.CreateTable())
}
