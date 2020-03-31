package main

import (
	"github.com/JasonSteinberg/timeTicker/database"
)

func main() {
	database.LoadDatabaseConfig("./production.json")
}
