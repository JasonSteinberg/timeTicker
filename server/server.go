package server

import (
	// "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JasonSteinberg/timeTicker/structs"
	"github.com/JasonSteinberg/timeTicker/users"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var ultraSecret = "buodcx3d4t06f0m1ld89ABCDEFGHIJKLMNOPQRSTUVWXYZfqpls" // Change and Do *NOT* put on github

func SetUpApi() {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", healthcheck).Methods("GET")
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", ProtectedMiddleWare(protected)).Methods("GET")

	log.Println("Starting server on port 8808.")
	log.Fatal(http.ListenAndServe(":8808", router)) // <- Do *NOT* use unencrypted version in production
	// log.Fatal(http.ListenAndServeTLS(":8808", "certificate.pem", "key.pem", router))  // <-- Use this one! (After you generate encryption)
}

// curl -v http://localhost:8808/healthcheck
func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"Status":"Good"}`)
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

	tokenString, err := token.SignedString([]byte(ultraSecret))

	if err != nil {
		log.Fatal("Oh no token failure! ", err)
	}

	return tokenString
}

// curl -X POST -d '{"email":"count_dooku","password":"iHateWookies"}' http://localhost:8808/signup
func signup(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	var error structs.ServerMessage

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing."
		ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hash)
	err = users.CreateUser(user)

	if err != nil {
		error.Message = err.Error()
		ErrorResponder(w, http.StatusInternalServerError, error)
		return
	}

	HappyMessage := structs.ServerMessage{Message: "Created user successfully"}
	w.Header().Set("Content-Type", "application/json")
	Responder(w, HappyMessage)
}

// curl -X POST -d '{"email":"count_dooku","password":"iHateWookies"}' http://localhost:8808/login
func login(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	var jwToken structs.JWT
	var error structs.ServerMessage

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing."
		ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		ErrorResponder(w, http.StatusBadRequest, error)
		return
	}

	password := user.Password
	login, err := users.CheckUser(user, password)

	if err != nil || login == false {
		error.Message = err.Error()
		ErrorResponder(w, http.StatusUnauthorized, error)
		return
	}

	token := GenerateToken(user)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	jwToken.Token = token

	Responder(w, jwToken)
}

// curl -X GET --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImNvdW50X2Rvb2t1IiwiaXNzIjoiY291cnNlIn0.osrQe3VwnTGqjuhHg36R9DRDt5apXSqb5-5CltMdp6g" http://localhost:8808/protected
func protected(w http.ResponseWriter, r *http.Request) {
	HappyMessage := structs.ServerMessage{Message: "You now have access to the protected route!"}
	w.Header().Set("Content-Type", "application/json")
	Responder(w, HappyMessage)
}

func ProtectedMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorObject structs.ServerMessage
		authToken := r.Header.Get("Authorization")

		token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("There was an error!")
			}
			return []byte(ultraSecret), nil
		})

		if error != nil {
			errorObject.Message = error.Error()
			ErrorResponder(w, http.StatusUnauthorized, errorObject)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			errorObject.Message = error.Error()
			ErrorResponder(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}
