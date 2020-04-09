package time

import (
	"github.com/JasonSteinberg/timeTicker/middleware"
	"github.com/gorilla/mux"
)

func SetUpTaskRoutes(router *mux.Router) {
	router.HandleFunc("/task/{id}/add", middleware.ProtectedMiddleWare(timeAdd)).Methods("POST")
}
