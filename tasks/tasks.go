package tasks

import (
	"github.com/JasonSteinberg/timeTicker/middleware"
	"github.com/gorilla/mux"
)

func SetUpTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", middleware.ProtectedMiddleWare(tasks)).Methods("GET")
	router.HandleFunc("/task/completed", middleware.ProtectedMiddleWare(taskCompleted)).Methods("GET")
	router.HandleFunc("/task/", middleware.ProtectedMiddleWare(taskNew)).Methods("POST")
	router.HandleFunc("/task/{id}", middleware.ProtectedMiddleWare(task)).Methods("GET")
	router.HandleFunc("/task/{id}", middleware.ProtectedMiddleWare(taskDelete)).Methods("DELETE")
}
