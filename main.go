package main

import (
	"github.com/JasonSteinberg/timeTicker/database"
	"github.com/JasonSteinberg/timeTicker/server"
	"github.com/JasonSteinberg/timeTicker/utils"
)

func main() {
	database.LoadDatabaseConfig("./production.json")
	utils.CheckForEmptyDatabase()
	server.SetUpApi()

}
