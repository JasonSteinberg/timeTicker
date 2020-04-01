package structs

import "time"

type ServerMessage struct {
	Message string `json:"message"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   int    `json:"access"`
}

type Task struct {
	ID          int       `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	DueDate     time.Time `json:"due_date"`
	IsCompleted bool      `json:"is_completed"`
}

type JWT struct {
	Token string `json:"token"`
}

var UltraSecret = "buodcx3d4t06f0m1ld89ABCDEFGHIJKLMNOPQRSTUVWXYZfqpls" // Change and Do *NOT* put on github
