package structs

type DatabaseConfig struct {
	User   	   string `json:"user"`
	Password   string `json:"password"`
	Url    	   string `json:"url"`
	Port   	   string `json:"port"`
	DbName 	   string `json:"dbname"`
}

var Database DatabaseConfig
