package database

import (
	"encoding/json"
	"github.com/JasonSteinberg/timeTicker/structs"
	"log"
	"os"
)

func LoadDatabaseConfig(fileName string) {
	configFile, err := os.Open(fileName)
	defer configFile.Close()

	settings := structs.DatabaseConfig{}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&settings)

	if err != nil {
		log.Fatal(`Fatal Error: Unable to load database configuration!
						Looking for: `, fileName)
	}

	structs.Database = settings
}
