package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func SetUpApi() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck).Methods("GET")

	log.Println("Starting server on port 8808.")
	log.Fatal(http.ListenAndServe(":8808", router))

}

func healthcheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"Status":"Good"}`)
}
