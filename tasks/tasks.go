package tasks

import (
	"github.com/JasonSteinberg/timeTicker/middleware"
	"github.com/gorilla/mux"
)

func SetUpTaskRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", middleware.ProtectedMiddleWare(tasks)).Methods("GET")
	router.HandleFunc("/task/{id}", middleware.ProtectedMiddleWare(task)).Methods("GET")
}
