package utils

import (
	"github.com/JasonSteinberg/timeTicker/database"
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
		setupDatabaseTables()
	}
}

func setupDatabaseTables() {

}