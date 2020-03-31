package structs

type ServerMessage struct {
	Message string `json:"message"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   int    `json:"access"`
}

type JWT struct {
	Token string `json:"token"`
}
