package database

import (
	"database/sql"
	"encoding/json"
	"github.com/JasonSteinberg/timeTicker/structs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func LoadDatabaseConfig(fileName string) {
	configFile, err := os.Open(fileName)
	defer configFile.Close()

	settings := structs.DatabaseConfig{}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&settings)

	if err != nil {
		log.Fatal(`Fatal Error: Unable to load database configuration!
						Looking for file : `, fileName)
	}

	structs.Database = settings
}

func GetConnectionString() string {
	var connectionString strings.Builder
	connectionString.WriteString(structs.Database.User)
	connectionString.WriteString(":")
	connectionString.WriteString(structs.Database.Password)
	connectionString.WriteString("@tcp(")
	connectionString.WriteString(structs.Database.Url)
	connectionString.WriteString(":")
	connectionString.WriteString(structs.Database.Port)
	connectionString.WriteString(")/")
	connectionString.WriteString(structs.Database.DbName)
	connectionString.WriteString("?parseTime=true")
	return connectionString.String()
}

func getSqlDB() *sql.DB {
	openString := GetConnectionString()

	db, err := sql.Open("mysql", openString)
	if err != nil {
		log.Fatal(`Fatal Error: Unable to connect to database, check configuration file! Error:`, err)
	}

	return db
}

func GetSqlReadDB() *sql.DB {
	return getSqlDB()
}

func GetSqlWriteDB() *sql.DB {
	return getSqlDB()
}

// From https://stackoverflow.com/questions/19991541
func ReturnJson(rows *sql.Rows) (string, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	tableData := make([]map[string]interface{}, 0)

	count := len(columns)
	values := make([]interface{}, count)
	scanArgs := make([]interface{}, count)
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(scanArgs...)
		if err != nil {
			return "", err
		}

		entry := make(map[string]interface{})
		for i, col := range columns {
			v := values[i]

			b, ok := v.([]byte)
			if ok {
				entry[col] = string(b)
			} else {
				entry[col] = v
			}
		}

		tableData = append(tableData, entry)
	}

	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func NewNullString(s string) sql.NullString {
	if len(s) == 0 || s == "<nil>" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewNullDate(t time.Time) sql.NullTime {
	if strings.HasPrefix(t.String(), "0001-01-01") {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func NewNullInt(s string) sql.NullInt64 {
	if len(s) == 0 || s == "<nil>" {
		return sql.NullInt64{
			Valid: false,
		}
	}

	tmpInt, err := strconv.Atoi(s)
	if err != nil {
		return sql.NullInt64{
			Valid: false,
		}
	}

	return sql.NullInt64{
		Int64: int64(tmpInt),
		Valid: true,
	}
}
