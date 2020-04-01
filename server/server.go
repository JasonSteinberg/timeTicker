package server

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"github.com/JasonSteinberg/timeTicker/middleware"
	"github.com/JasonSteinberg/timeTicker/structs"
	"github.com/JasonSteinberg/timeTicker/tasks"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func SetUpApi() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck).Methods("GET")
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", middleware.ProtectedMiddleWare(protected)).Methods("GET")

	tasks.SetUpTaskRoutes(router)


	log.Println("Starting server on port 8808.")
	log.Fatal(http.ListenAndServe(":8808", router)) // <- Do *NOT* use unencrypted version in production
	// log.Fatal(http.ListenAndServeTLS(":8808", "certificate.pem", "key.pem", router))  // <-- Use this one! (After you generate encryption)
}

// curl -v http://localhost:8808/healthcheck
func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"Status":"Good"}`)
}

// curl -X GET --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNvdW50X2Rvb2t1IiwiaXNzIjoiY291cnNlIn0.osrQe3VwnTGqjuhHg36R9DRDt5apXSqb5-5CltMdp6g" http://localhost:8808/protected
func protected(w http.ResponseWriter, r *http.Request) {
	HappyMessage := structs.ServerMessage{Message: "You now have access to the protected route!"}
	w.Header().Set("Content-Type", "application/json")
	Responder(w, HappyMessage)
}

func ErrorResponder(w http.ResponseWriter, status int, error structs.ServerMessage) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func Responder(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}

func GenerateToken(user structs.User) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(structs.UltraSecret))

	if err != nil {
		log.Fatal("Oh no token failure! ", err)
	}

	return tokenString
}
